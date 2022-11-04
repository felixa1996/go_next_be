package domain_user_dto

type UserDtoCreateInput struct {
	Name   string `validate:"required,min=4,max=225"`
	Author string `validate:"required,min=4,max=225"`
}

type UserDtoCreateOutput struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Author string `json:"author"`
}

type UserDtoDeleteInput struct {
	Id string `validate:"required"`
}

type UserDtoFindOneInput struct {
	Id string `validate:"required"`
}

type UserDtoUpdateParamInput struct {
	Id string `validate:"required"`
}
type UserDtoUpdateInput struct {
	Name   string `validate:"required,min=4,max=15"`
	Author string `validate:"required,min=4,max=15"`
}
