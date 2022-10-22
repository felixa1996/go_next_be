package domain_company_handler

import (
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	domain "github.com/felixa1996/go_next_be/app/domain/company"
	repository "github.com/felixa1996/go_next_be/app/domain/company/repository"
	usecase "github.com/felixa1996/go_next_be/app/domain/company/usecase"
	"github.com/felixa1996/go_next_be/app/infra/database"
)

type CompanyHandler struct {
	usecase    domain.CompanyUsecaseContract
	logger     *zap.Logger
	validate   *validator.Validate
	translator ut.Translator
	sqs        *sqs.SQS
}

func RegisterCompanyEventHandler(db database.Manager, sqs *sqs.SQS, logger *zap.Logger, validate *validator.Validate, translator ut.Translator, contextTimeout time.Duration) {
	// init event handler
	repo := repository.NewCompanyMongoRepository(&db, logger)
	usecase := usecase.NewCompanyUsecase(repo, logger, contextTimeout)
	handler := &CompanyHandler{
		usecase:    usecase,
		logger:     logger,
		validate:   validate,
		translator: translator,
		sqs:        sqs,
	}

	handler.Upsert("a")
}

// todo add delete sqs here
