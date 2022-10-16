package domain_user_usecase

import (
	"context"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (u *userUsecase) Create(c context.Context, dto dto.UserDtoCreateInput) (domain.User, error) {
	ctx := context.TODO()

	user := domain.User{
		Id:     uuid.NewString(),
		Name:   dto.Name,
		Author: dto.Author,
	}

	res, err := u.repo.Create(ctx, user)
	if err != nil {
		u.logger.Error("Failed to create user usecase", zap.Error(err))
		return domain.User{}, err
	}
	return res, nil
}
