package domain_media

import (
	"context"

	dto "github.com/felixa1996/go_next_be/app/domain/media/dto"
)

const CollectionName = "media"

type Media struct {
	Id         string `bson:"id,omitempty" json:"id"`
	Uri        string `bson:"uri" json:"uri"`
	Curicullum string `bson:"curicullum" json:"curicullum"`
	Year       string `bson:"year" json:"year"`
	Subject    string `bson:"subject" json:"subject"`
}

// MediaUsecaseContract repersent usecase contract tied with media domain
type MediaUsecaseContract interface {
	Create(context.Context, dto.MediaDtoCreateInput) (Media, error)
}

// MediaRepository represent repository contract
type MediaRepository interface {
	Create(context.Context, Media) (Media, error)
}
