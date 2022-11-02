package domain_media_handler

import (
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	domain "github.com/felixa1996/go_next_be/app/domain/media"
	repository "github.com/felixa1996/go_next_be/app/domain/media/repository"
	usecase "github.com/felixa1996/go_next_be/app/domain/media/usecase"
	"github.com/felixa1996/go_next_be/app/infra/database"
	"github.com/felixa1996/go_next_be/app/infra/uploader"
)

type MediaHandler struct {
	minioWrapper uploader.MinioWrapper
	usecase      domain.MediaUsecaseContract
	logger       *zap.Logger
	validate     *validator.Validate
	translator   ut.Translator
}

func RegisterMediaHandler(db database.Manager, minioWrapper uploader.MinioWrapper, logger *zap.Logger, validate *validator.Validate, translator ut.Translator, contextTimeout time.Duration, group *echo.Group) {
	// init handler
	repo := repository.NewMediaMongoRepository(&db, logger)
	usecase := usecase.NewMediaUsecase(minioWrapper, repo, logger, contextTimeout)
	handler := &MediaHandler{
		// todo need to remove
		minioWrapper: minioWrapper,
		usecase:      usecase,
		logger:       logger,
		validate:     validate,
		translator:   translator,
	}

	group.POST("", handler.Create)
}
