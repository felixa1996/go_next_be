package domain_user_repository

import (
	"context"
	"net/http"

	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"

	. "github.com/felixa1996/go_next_be/app/infra/error"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
)

func (r *userMongoRepository) Update(ctx context.Context, user domain.User) (domain.User, error) {
	_, err := r.db.Database.Collection(domain.CollectionName).UpdateOne(context.TODO(), bson.M{"id": user.Id}, user)
	if err != nil {
		r.logger.Error("Failed to update user repository", zap.Error(err))
		return domain.User{}, NewErrorWrapper(http.StatusInternalServerError, err, "Failed to update user repository")
	}

	return user, nil
}
