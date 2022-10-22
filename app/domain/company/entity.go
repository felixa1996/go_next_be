package domain_company

import (
	"context"

	dto "github.com/felixa1996/go_next_be/app/domain/company/dto"
)

const (
	CollectionName = "company"
	// todo need to env
	UpsertQueueUrl = "https://sqs.us-east-1.amazonaws.com/065561208089/test_sqs_data.fifo"
)

type Company struct {
	Id          string `bson:"id,omitempty" json:"id"`
	CompanyName string `bson:"company_name" json:"company_name"`
}

// CompanyUsecaseContract repersent usecase contract tied with user domain
type CompanyUsecaseContract interface {
	// todo need remove company
	Upsert(context.Context, dto.CompanyDtoUpsert) (Company, error)
}

// CompanyRepository represent repository contract
type CompanyRepository interface {
	Upsert(context.Context, Company) (Company, error)
}
