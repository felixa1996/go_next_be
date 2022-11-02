package domain_media_dto

import "mime/multipart"

type MediaDtoCreateInput struct {
	Uri   string `validate:"required,min=4,max=15" json:"uri"`
	Files []*multipart.FileHeader
}

type MediaDtoCreateOutput struct {
	Id  string `json:"id"`
	Uri string `json:"uri"`
}
