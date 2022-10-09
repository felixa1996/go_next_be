package domain_user_repository

import (
	domain "github.com/felixa1996/go_next_be/app/domain/user"
	"github.com/felixa1996/go_next_be/app/infra/database"
	"go.uber.org/zap"
)

type userMongoRepository struct {
	db     *database.Manager
	logger *zap.Logger
}

func NewUserMongoRepository(db *database.Manager, logger *zap.Logger) domain.UserRepository {
	return &userMongoRepository{
		db:     db,
		logger: logger,
	}
}
