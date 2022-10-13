package domain_user_usecase

import (
	"context"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
	"go.uber.org/zap"
)

func (u *userUsecase) Update(c context.Context, param dto.UserDtoUpdateParamInput, dto dto.UserDtoUpdateInput) (domain.User, error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	user := domain.User{
		Id:     param.Id,
		Name:   dto.Name,
		Author: dto.Author,
	}

	res, err := u.repo.Update(ctx, user)
	if err != nil {
		u.logger.Error("Failed to update user usecase", zap.Error(err))
		return domain.User{}, err
	}
	return res, nil
}
