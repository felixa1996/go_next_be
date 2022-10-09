package domain_user

import (
	"context"
	"log"

	"github.com/felixa1996/go_next_be/app/infra/database"
	"gopkg.in/mgo.v2/bson"
)

type userMongoRepository struct {
	Db *database.Manager
}

func NewUserMongoRepository(db *database.Manager) UserRepository {
	return &userMongoRepository{
		Db: db,
	}
}

func (r *userMongoRepository) FindPagination(ctx context.Context) ([]User, error) {
	var user User
	var users []User

	// todo need change default db
	cursor, err := r.Db.Client.Database("generic_db").Collection(CollectionName).Find(ctx, bson.M{})
	if err != nil {
		log.Fatal("Failed to fetch user", err)
	}
	for cursor.Next(ctx) {
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	defer cursor.Close(context.TODO())

	return users, nil
}
