package domain_company_dto

type CompanyDtoUpsert struct {
	Id          string `json:"id"`
	CompanyName string `validate:"required,min=4,max=15" json:"company_name"`
}
