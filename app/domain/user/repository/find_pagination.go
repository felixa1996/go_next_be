package domain_user_repository

import (
	"context"
	"net/http"

	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"

	. "github.com/felixa1996/go_next_be/app/infra/error"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
)

func (r *userMongoRepository) FindPagination(ctx context.Context) ([]domain.User, error) {
	var user domain.User
	var users []domain.User

	cursor, err := r.db.Database.Collection(domain.CollectionName).Find(ctx, bson.M{})
	if err != nil {
		r.logger.Error("Failed to fetch user repository", zap.Error(err))
		return []domain.User{}, NewErrorWrapper(http.StatusInternalServerError, err, "Failed to fetch user repository")
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&user)
		if err != nil {
			r.logger.Error("Failed to decode user repository", zap.Error(err))
			return []domain.User{}, NewErrorWrapper(http.StatusInternalServerError, err, "Failed to decode user repository")
		}
		users = append(users, user)
	}
	defer cursor.Close(context.TODO())

	return users, nil
}
