package domain_user

import (
	"context"

	"github.com/felixa1996/go_next_be/app/infra/database"
	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

type userMongoRepository struct {
	db     *database.Manager
	logger *zap.Logger
}

func NewUserMongoRepository(db *database.Manager, logger *zap.Logger) UserRepository {
	return &userMongoRepository{
		db:     db,
		logger: logger,
	}
}

func (r *userMongoRepository) FindPagination(ctx context.Context) ([]User, error) {
	var user User
	var users []User

	// todo need change default db
	cursor, err := r.db.Client.Database("generic_db").Collection(CollectionName).Find(ctx, bson.M{})
	if err != nil {
		r.logger.Fatal("Failed to fetch user", zap.String("err", err.Error()))
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&user)
		if err != nil {
			r.logger.Fatal("Failed to decode user", zap.String("err", err.Error()))
		}
		users = append(users, user)
	}
	defer cursor.Close(context.TODO())

	return users, nil
}
