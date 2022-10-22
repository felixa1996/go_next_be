package domain_company_repository

import (
	domain "github.com/felixa1996/go_next_be/app/domain/company"
	"github.com/felixa1996/go_next_be/app/infra/database"
	"go.uber.org/zap"
)

type companyMongoRepository struct {
	db     *database.Manager
	logger *zap.Logger
}

func NewCompanyMongoRepository(db *database.Manager, logger *zap.Logger) domain.CompanyRepository {
	return &companyMongoRepository{
		db:     db,
		logger: logger,
	}
}
