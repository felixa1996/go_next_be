package domain_user

import (
	"context"

	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
)

const CollectionName = "user"

type User struct {
	Id     string `bson:"id" json:"id"`
	Name   string `bson:"name" json:"name"`
	Author string `bson:"author,omitempty" json:"author"`
}

// UserUsecaseContract repersent usecase contract tied with user domain
type UserUsecaseContract interface {
	FindPagination(context.Context) ([]User, error)
	Create(context.Context, dto.UserDtoCreateInput) (User, error)
}

// UserRepository represent repository contract
type UserRepository interface {
	FindPagination(context.Context) ([]User, error)
	Create(context.Context, User) (User, error)
}
