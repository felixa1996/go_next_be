package domain_user

import "context"

const CollectionName = "user"

type User struct {
	Name   string `bson:"name" json:"name"`
	Author string `bson:"author,omitempty" json:"author"`
}

// UserUsecaseContract repersent usecase contract tied with user domain
type UserUsecaseContract interface {
	FindPagination(context.Context) ([]User, error)
}

// UserRepository represent repository contract
type UserRepository interface {
	FindPagination(context.Context) ([]User, error)
}
