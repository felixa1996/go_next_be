package domain_media_repository

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	. "github.com/felixa1996/go_next_be/app/infra/error"

	domain "github.com/felixa1996/go_next_be/app/domain/media"
)

func (r *mediaMongoRepository) Create(ctx context.Context, media domain.Media) (domain.Media, error) {
	_, err := r.db.Database.Collection(domain.CollectionName).InsertOne(context.TODO(), media)
	if err != nil {
		r.logger.Error("Failed to create media repository", zap.Error(err))
		return domain.Media{}, NewErrorWrapper(http.StatusInternalServerError, err, "Failed to create media repository")
	}

	return media, nil
}
