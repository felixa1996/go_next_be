package domain_user_repository

import (
	"context"
	"errors"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	"github.com/felixa1996/go_next_be/app/infra/database"
	mocks "github.com/felixa1996/go_next_be/mocks/app/domain/user"
)

type testStruct struct {
	Name                 string
	Message              string
	Data                 domain.User
	DataError            error
	CreateResponse       interface{}
	ExpectSuccessReponse domain.User
	ExpectErrorReponse   error
}

func setupTestEnv(t *testing.T) (faker.Faker, *zap.Logger) {
	fake := faker.New()
	logger := zaptest.NewLogger(t)
	return fake, logger
}

func TestUserRepositoryCreate(t *testing.T) {
	t.Parallel()
	fake, logger := setupTestEnv(t)

	data := domain.User{
		Id:     fake.UUID().V4(),
		Name:   fake.Person().Name(),
		Author: fake.Person().FirstName(),
	}

	repoStruct := []testStruct{
		{
			Name:                 "Success",
			Message:              "should success",
			Data:                 data,
			CreateResponse:       data,
			ExpectSuccessReponse: data,
		},
		{
			Name:               "Failed",
			Message:            "should failed",
			Data:               data,
			DataError:          errors.New("failed to create user"),
			ExpectErrorReponse: errors.New("failed to create user"),
		},
	}

	for _, tc := range repoStruct {
		mockRepository := new(mocks.UserRepository)
		mockRepository.On("Create", tc.Data).Return(tc.CreateResponse)

		mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		defer mt.Close()

		mt.Run(tc.Name, func(t *mtest.T) {
			if tc.ExpectErrorReponse != nil {
				t.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
					Message: tc.DataError.Error(),
				}))
			} else {
				t.AddMockResponses(mtest.CreateSuccessResponse(
					primitive.E{Key: "Id", Value: tc.Data.Id},
					primitive.E{Key: "Name", Value: tc.Data.Name},
					primitive.E{Key: "Author", Value: tc.Data.Author},
				))
			}

			databaseManager := &database.Manager{
				Database: t.DB,
			}
			repo := NewUserMongoRepository(databaseManager, logger)
			res, err := repo.Create(context.TODO(), tc.Data)

			if err != nil {
				assert.Contains(t, err.Error(), tc.ExpectErrorReponse.Error(), tc.Message)
				return
			}
			assert.Equal(t, tc.ExpectSuccessReponse, res, tc.Message)
		})
	}
}
