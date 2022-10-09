package domain_user_usecase

import (
	"context"

	"github.com/google/uuid"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
	"go.uber.org/zap"
)

func (u *userUsecase) Create(c context.Context, dto dto.UserDtoCreateInput) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user := domain.User{
		Id:     uuid.NewString(),
		Name:   dto.Name,
		Author: dto.Author,
	}

	res, err := u.repo.Create(ctx, user)
	if err != nil {
		u.logger.Error("Failed to create user", zap.Error(err))
		return domain.User{}, err
	}
	return res, nil
}
