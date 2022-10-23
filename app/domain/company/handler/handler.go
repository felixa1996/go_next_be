package domain_company_handler

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	"github.com/felixa1996/go_next_be/app/config"
	domain "github.com/felixa1996/go_next_be/app/domain/company"
	repository "github.com/felixa1996/go_next_be/app/domain/company/repository"
	usecase "github.com/felixa1996/go_next_be/app/domain/company/usecase"
	"github.com/felixa1996/go_next_be/app/infra/database"
	message "github.com/felixa1996/go_next_be/app/infra/message"
)

type CompanyHandler struct {
	config     *config.Config
	usecase    domain.CompanyUsecaseContract
	logger     *zap.Logger
	validate   *validator.Validate
	translator ut.Translator
	sqs        *sqs.SQS
}

func RegisterCompanyEventHandler(config *config.Config, db database.Manager, sqs *sqs.SQS, logger *zap.Logger, validate *validator.Validate, translator ut.Translator, contextTimeout time.Duration) {
	// init event handler
	repo := repository.NewCompanyMongoRepository(&db, logger)
	usecase := usecase.NewCompanyUsecase(repo, logger, contextTimeout)
	handler := &CompanyHandler{
		config:     config,
		usecase:    usecase,
		logger:     logger,
		validate:   validate,
		translator: translator,
		sqs:        sqs,
	}

	handler.Upsert()
}

// Delete Message place here because it's geeric
func (h *CompanyHandler) deleteMessage(msg message.SqsIncomingMessage) {
	_, err := h.sqs.DeleteMessage(&sqs.DeleteMessageInput{
		QueueUrl:      aws.String(msg.QueueUrl),
		ReceiptHandle: msg.Message.ReceiptHandle,
	})
	if err != nil {
		h.logger.Error("Failed to delete sqs message",
			zap.String("QueueUrl", msg.QueueUrl),
			zap.String("MessageId", *msg.Message.MessageId),
			zap.String("ReceiptHandle", *msg.Message.ReceiptHandle),
			zap.String("MessageBody", *msg.Message.Body),
			zap.Error(err),
		)
	}
	h.logger.Info("Success to delete sqs message", zap.String("QueueUrl", msg.QueueUrl),
		zap.String("MessageId", *msg.Message.MessageId),
		zap.String("ReceiptHandle", *msg.Message.ReceiptHandle),
		zap.String("MessageBody", *msg.Message.Body),
	)
}
