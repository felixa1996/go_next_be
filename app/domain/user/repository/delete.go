package domain_user_repository

import (
	"context"
	"net/http"

	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"

	. "github.com/felixa1996/go_next_be/app/infra/error"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
)

func (r *userMongoRepository) Delete(ctx context.Context, id string) error {
	// check if user exists
	_, err := r.FindOne(context.TODO(), id)
	if err != nil {
		return err
	}

	_, err = r.db.Database.Collection(domain.CollectionName).DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		r.logger.Error("Failed to delete user repository", zap.Error(err))
		return NewErrorWrapper(http.StatusInternalServerError, err, "Failed to delete user repository")
	}

	return nil
}
