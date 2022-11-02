package domain_media_usecase

import (
	"time"

	"go.uber.org/zap"

	domain "github.com/felixa1996/go_next_be/app/domain/media"
	"github.com/felixa1996/go_next_be/app/infra/uploader"
)

type mediaUsecase struct {
	minioWrapper   uploader.MinioWrapper
	repo           domain.MediaRepository
	contextTimeout time.Duration
	logger         *zap.Logger
}

func NewMediaUsecase(minioWrapper uploader.MinioWrapper, repo domain.MediaRepository, logger *zap.Logger, contextTimeout time.Duration) domain.MediaUsecaseContract {
	return &mediaUsecase{
		minioWrapper:   minioWrapper,
		repo:           repo,
		contextTimeout: contextTimeout,
		logger:         logger,
	}
}
