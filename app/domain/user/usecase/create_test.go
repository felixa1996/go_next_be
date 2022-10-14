package domain_user_usecase

import (
	"context"
	"testing"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	domain_user_dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
	repo "github.com/felixa1996/go_next_be/app/domain/user/repository"
	"github.com/felixa1996/go_next_be/app/infra/database"
	mocks "github.com/felixa1996/go_next_be/mocks/app/domain/user"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.uber.org/zap/zaptest"
)

func TestUserCreate(t *testing.T) {
	logger := zaptest.NewLogger(t)

	args2 := &domain.User{
		Id:     "sdsdsdsd",
		Name:   "test",
		Author: "sdsdsd",
	}

	mockRepository := new(mocks.UserRepository)
	mockRepository.On("Create", args2).Return(args2)

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mockUsecase := new(mocks.UserUsecaseContract)
	mockUsecase.On("Create", domain_user_dto.UserDtoCreateInput{
		Name: "test",
	}).Return(args2)

	mt.Run("Success", func(t *mtest.T) {
		databaseManager := &database.Manager{
			Database: t.DB,
		}
		repo := repo.NewUserMongoRepository(databaseManager, logger)
		usecase := NewUserUsecase(repo, logger, 100000000000)
		usecase.Create(context.TODO(), domain_user_dto.UserDtoCreateInput{
			Name:   "test",
			Author: "test",
		})
	})

	// if err != nil {
	// 	println(err.Error())
	// }
	// print(res.Author)
}
