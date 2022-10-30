package domain_media_dto

type MediaDtoCreateInput struct {
	Uri string `validate:"required,min=4,max=15"`
}

type MediaDtoCreateOutput struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}
