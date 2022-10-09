package domain_user

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type userUsecase struct {
	repo           UserRepository
	contextTimeout time.Duration
	logger         *zap.Logger
}

func NewUserUsecase(repo UserRepository, logger *zap.Logger, contextTimeout time.Duration) UserUsecaseContract {
	return &userUsecase{
		repo:           repo,
		contextTimeout: contextTimeout,
		logger:         logger,
	}
}

func (u *userUsecase) FindPagination(c context.Context) ([]User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.repo.FindPagination(ctx)
	if err != nil {
		u.logger.Fatal("Failed to fetch user", zap.Error(err))
		return nil, err
	}
	return res, nil
}
