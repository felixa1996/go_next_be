package domain_company_handler

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	domain_company "github.com/felixa1996/go_next_be/app/domain/company"
	dto "github.com/felixa1996/go_next_be/app/domain/company/dto"
	"go.uber.org/zap"
)

func (h *CompanyHandler) Upsert(message string) {
	h.logger.Info("Incoming sqs", zap.String("message", message))

	chnMessages := make(chan *sqs.Message, 2)
	go h.pollMessages(chnMessages)

	for msg := range chnMessages {
		err := h.handleMessage(msg)
		if err != nil {
			continue
		}
		h.deleteMessage(msg)
	}
}

// todo change parameter to struct
func (h *CompanyHandler) handleMessage(msg *sqs.Message) error {
	data := &domain_company.Company{}
	err := json.Unmarshal([]byte(*msg.Body), &data)
	if err != nil {
		h.logger.Error("Failed to decode message body company",
			zap.String("QueueUrl", domain_company.UpsertQueueUrl),
			zap.String("MessageId", *msg.MessageId),
			zap.String("ReceiptHandle", *msg.ReceiptHandle),
			zap.String("MessageBody", *msg.Body),
			zap.Error(err))
		return err
	}

	dto := dto.CompanyDtoUpsert{
		Id:          data.Id,
		CompanyName: data.CompanyName,
	}

	err = h.validate.Struct(dto)
	if err != nil {
		h.logger.Error("Failed to validate upsert company",
			zap.String("QueueUrl", domain_company.UpsertQueueUrl),
			zap.String("MessageId", *msg.MessageId),
			zap.String("ReceiptHandle", *msg.ReceiptHandle),
			zap.String("MessageBody", *msg.Body),
			zap.Error(err))
		return err
	}
	_, err = h.usecase.Upsert(context.TODO(), dto)
	if err != nil {
		return err
	}
	return nil
}

func (h *CompanyHandler) deleteMessage(msg *sqs.Message) {
	_, err := h.sqs.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(domain_company.UpsertQueueUrl),
		ReceiptHandle: msg.ReceiptHandle,
	})
	if err != nil {
		h.logger.Error("Failed to acknowledge sqs message",
			zap.String("QueueUrl", domain_company.UpsertQueueUrl),
			zap.String("MessageId", *msg.MessageId),
			zap.String("ReceiptHandle", *msg.ReceiptHandle),
			zap.String("MessageBody", *msg.Body),
			zap.Error(err),
		)
	}
	h.logger.Info("Success to acknowledge sqs message", zap.String("QueueUrl", domain_company.UpsertQueueUrl),
		zap.String("MessageId", *msg.MessageId),
		zap.String("ReceiptHandle", *msg.ReceiptHandle),
		zap.String("MessageBody", *msg.Body),
	)
}

func (h *CompanyHandler) pollMessages(chn chan<- *sqs.Message) {

	for {
		output, err := h.sqs.ReceiveMessage(&sqs.ReceiveMessageInput{
			QueueUrl: aws.String(domain_company.UpsertQueueUrl),
			// todo change to env
			MaxNumberOfMessages: aws.Int64(2),
			WaitTimeSeconds:     aws.Int64(15),
		})

		if err != nil {
			h.logger.Error("Failed to fetch sqs message",
				zap.String("QueueUrl", domain_company.UpsertQueueUrl),
				zap.Error(err),
			)
		}

		for _, message := range output.Messages {
			chn <- message
		}

	}

}
