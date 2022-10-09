package domain_user_repository

import (
	"context"

	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
)

func (r *userMongoRepository) FindPagination(ctx context.Context) ([]domain.User, error) {
	var user domain.User
	var users []domain.User

	// todo need change default db
	cursor, err := r.db.Client.Database("generic_db").Collection(domain.CollectionName).Find(ctx, bson.M{})
	if err != nil {
		r.logger.Fatal("Failed to fetch user", zap.Error(err))
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&user)
		if err != nil {
			r.logger.Fatal("Failed to decode user", zap.Error(err))
		}
		users = append(users, user)
	}
	defer cursor.Close(context.TODO())

	return users, nil
}
