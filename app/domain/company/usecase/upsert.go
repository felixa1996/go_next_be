package domain_company_usecase

import (
	"context"

	domain "github.com/felixa1996/go_next_be/app/domain/company"
	dto "github.com/felixa1996/go_next_be/app/domain/company/dto"
	"go.uber.org/zap"
)

func (u *companyUsecase) Upsert(c context.Context, dto dto.CompanyDtoUpsert) error {
	ctx := context.TODO()

	u.logger.Info("Processing upsert company usecase", zap.Any("Object", dto))
	company := domain.Company{
		Id:          dto.Id,
		CompanyName: dto.CompanyName,
	}

	_, err := u.repo.Upsert(ctx, company)
	if err != nil {
		u.logger.Error("Failed to upsert company usecase", zap.Error(err))
		return err
	}
	u.logger.Info("Success upsert company usecase", zap.Any("Object", dto))
	return nil
}
