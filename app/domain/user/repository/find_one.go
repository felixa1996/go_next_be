package domain_user_repository

import (
	"context"
	"net/http"

	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"

	. "github.com/felixa1996/go_next_be/app/infra/error"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
)

func (r *userMongoRepository) FindOne(ctx context.Context, id string) (domain.User, error) {
	var user domain.User

	err := r.db.Database.Collection(domain.CollectionName).FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		r.logger.Error("User not found", zap.String("id", id), zap.Error(err))
		return domain.User{}, NewErrorWrapper(http.StatusInternalServerError, nil, "User not found")
	}

	return user, nil
}
