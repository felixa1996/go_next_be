package domain_company_usecase

import (
	"time"

	"go.uber.org/zap"

	domain "github.com/felixa1996/go_next_be/app/domain/company"
)

type companyUsecase struct {
	repo           domain.CompanyRepository
	contextTimeout time.Duration
	logger         *zap.Logger
}

func NewCompanyUsecase(repo domain.CompanyRepository, logger *zap.Logger, contextTimeout time.Duration) domain.CompanyUsecaseContract {
	return &companyUsecase{
		repo:           repo,
		contextTimeout: contextTimeout,
		logger:         logger,
	}
}
