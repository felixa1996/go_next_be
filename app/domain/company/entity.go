package domain_company

import (
	"context"

	dto "github.com/felixa1996/go_next_be/app/domain/company/dto"
)

const (
	CollectionName = "company"
)

type (
	Company struct {
		Id          string `bson:"id,omitempty" json:"id"`
		CompanyName string `bson:"company_name" json:"company_name"`
	}

	// CompanyUsecaseContract repersent usecase contract tied with user domain
	CompanyUsecaseContract interface {
		Upsert(context.Context, dto.CompanyDtoUpsert) error
	}

	// CompanyRepository represent repository contract
	CompanyRepository interface {
		Upsert(context.Context, Company) (Company, error)
	}
)
