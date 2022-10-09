package domain_user_usecase

import (
	"time"

	"go.uber.org/zap"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
)

type userUsecase struct {
	repo           domain.UserRepository
	contextTimeout time.Duration
	logger         *zap.Logger
}

func NewUserUsecase(repo domain.UserRepository, logger *zap.Logger, contextTimeout time.Duration) domain.UserUsecaseContract {
	return &userUsecase{
		repo:           repo,
		contextTimeout: contextTimeout,
		logger:         logger,
	}
}
