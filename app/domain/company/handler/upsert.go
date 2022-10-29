package domain_company_handler

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go/service/sqs"
	domain_company "github.com/felixa1996/go_next_be/app/domain/company"
	dto "github.com/felixa1996/go_next_be/app/domain/company/dto"
	message "github.com/felixa1996/go_next_be/app/infra/message"
	"go.uber.org/zap"
)

func (h *CompanyHandler) Upsert() {
	queueUrl := h.config.SqsCompanyUpsertUrl
	maxNumberMessage := h.config.SqsCompanyUpsertMaxNumberMessage
	waitTimeout := h.config.SqsCompanyWaitTimeOutSeconds

	chnMessages := make(chan *sqs.Message, 2)
	go h.sqsWrapper.PollMessages(queueUrl, maxNumberMessage, waitTimeout, chnMessages)

	for chnMessage := range chnMessages {
		msg := message.SqsIncomingMessage{
			QueueUrl: h.config.SqsCompanyUpsertUrl,
			Message:  chnMessage,
		}
		err := h.handleMessage(msg)
		if err != nil {
			continue
		}
		h.sqsWrapper.DeleteMessage(queueUrl, msg.Message)
	}
}

func (h *CompanyHandler) handleMessage(msg message.SqsIncomingMessage) error {
	data := &domain_company.Company{}
	err := json.Unmarshal([]byte(*msg.Message.Body), &data)
	if err != nil {
		h.logger.Error("Failed to decode message body company",
			zap.String("QueueUrl", h.config.SqsCompanyUpsertUrl),
			zap.String("MessageId", *msg.Message.MessageId),
			zap.String("ReceiptHandle", *msg.Message.ReceiptHandle),
			zap.String("MessageBody", *msg.Message.Body),
			zap.Error(err))
		h.sqsWrapper.DeleteMessage(msg.QueueUrl, msg.Message)
		return err
	}

	dto := dto.CompanyDtoUpsert{
		Id:          data.Id,
		CompanyName: data.CompanyName,
	}

	err = h.validate.Struct(dto)
	if err != nil {
		h.logger.Error("Failed to validate upsert company",
			zap.String("QueueUrl", msg.QueueUrl),
			zap.String("MessageId", *msg.Message.MessageId),
			zap.String("ReceiptHandle", *msg.Message.ReceiptHandle),
			zap.String("MessageBody", *msg.Message.Body),
			zap.Error(err))
		return err
	}

	err = h.usecase.Upsert(context.TODO(), dto)
	if err != nil {
		return err
	}
	return nil
}
