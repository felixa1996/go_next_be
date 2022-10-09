package domain_user

import (
	"context"
	"log"
	"time"
)

type userUsecase struct {
	repo           UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(repo UserRepository, contextTimeout time.Duration) UserUsecaseContract {
	return &userUsecase{
		repo:           repo,
		contextTimeout: contextTimeout,
	}
}

func (u *userUsecase) FindPagination(c context.Context) ([]User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.repo.FindPagination(ctx)
	if err != nil {
		log.Fatal("Failed to fetch user", err)
		return nil, err
	}
	return res, nil
}
