package domain_user_usecase

import (
	"context"

	"go.uber.org/zap"

	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
)

func (u *userUsecase) Delete(c context.Context, dto dto.UserDtoDeleteInput) error {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	err := u.repo.Delete(ctx, dto.Id)
	if err != nil {
		u.logger.Error("Failed to delete user usecase", zap.Error(err))
		return err
	}
	return nil
}
