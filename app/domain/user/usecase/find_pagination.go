package domain_user_usecase

import (
	"context"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	"go.uber.org/zap"
)

func (u *userUsecase) FindPagination(c context.Context) ([]domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	res, err := u.repo.FindPagination(ctx)
	if err != nil {
		u.logger.Fatal("Failed to fetch user", zap.Error(err))
		return nil, err
	}
	return res, nil
}
