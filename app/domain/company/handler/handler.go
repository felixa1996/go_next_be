package domain_company_handler

import (
	"time"

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
	sqsWrapper message.SqsWrapper
}

func RegisterCompanyEventHandler(config *config.Config, db database.Manager, sqsWrapper message.SqsWrapper, logger *zap.Logger, validate *validator.Validate, translator ut.Translator, contextTimeout time.Duration) {
	// init event handler
	repo := repository.NewCompanyMongoRepository(&db, logger)
	usecase := usecase.NewCompanyUsecase(repo, logger, contextTimeout)
	handler := &CompanyHandler{
		config:     config,
		usecase:    usecase,
		logger:     logger,
		validate:   validate,
		translator: translator,
		sqsWrapper: sqsWrapper,
	}

	handler.Upsert()
}
