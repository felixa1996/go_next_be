package domain_user_repository

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	. "github.com/felixa1996/go_next_be/app/infra/error"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
)

func (r *userMongoRepository) Create(ctx context.Context, user domain.User) (domain.User, error) {
	_, err := r.db.Database.Collection(domain.CollectionName).InsertOne(context.TODO(), user)
	if err == nil {
		r.logger.Error("Failed to create user repository", zap.Error(err))
		return domain.User{}, NewErrorWrapper(http.StatusInternalServerError, err, "Failed to create user repository")
	}

	return user, nil
}
