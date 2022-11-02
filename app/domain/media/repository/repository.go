package domain_media_repository

import (
	domain "github.com/felixa1996/go_next_be/app/domain/media"
	"github.com/felixa1996/go_next_be/app/infra/database"
	"go.uber.org/zap"
)

type mediaMongoRepository struct {
	db     *database.Manager
	logger *zap.Logger
}

func NewMediaMongoRepository(db *database.Manager, logger *zap.Logger) domain.MediaRepository {
	return &mediaMongoRepository{
		db:     db,
		logger: logger,
	}
}
