package domain_user_usecase

import (
	"context"

	"go.uber.org/zap"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
)

func (u *userUsecase) FindOne(c context.Context, dto dto.UserDtoFindOneInput) (domain.User, error) {
	ctx := context.TODO()

	res, err := u.repo.FindOne(ctx, dto.Id)
	if err != nil {
		u.logger.Error("Failed to find one user", zap.Error(err))
		return res, err
	}
	return res, nil
}
