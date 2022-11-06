package domain_user_repository

import (
	"context"
	"errors"
	"testing"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	"github.com/felixa1996/go_next_be/app/infra/database"
	"github.com/stretchr/testify/assert"
)

type testDeleteStruct struct {
	Name                      string
	Message                   string
	Data                      string
	DataError                 error
	DataFindOneError          error
	FindOneResponse           interface{}
	ExpectSuccessReponse      domain.User
	ExpectErrorReponse        error
	ExpectFindOneErrorReponse error
}

func TestUserDeleteRepository(t *testing.T) {
	t.Parallel()
	fake, logger := setupTestEnv(t)

	data := domain.User{
		Id:     fake.UUID().V4(),
		Name:   fake.Person().Name(),
		Author: fake.Person().FirstName(),
	}

	repoStruct := []testDeleteStruct{
		{
			Name:                      "Failed Findone",
			Message:                   "should failed findone",
			Data:                      data.Id,
			DataFindOneError:          errors.New("User not found"),
			ExpectFindOneErrorReponse: errors.New("User not found"),
		},
		{
			Name:               "Failed Delete",
			Message:            "should failed delete",
			Data:               data.Id,
			DataError:          errors.New("User not found"),
			ExpectErrorReponse: errors.New("User not found"),
		},
		{
			Name:                 "Success",
			Message:              "should success",
			Data:                 data.Id,
			ExpectSuccessReponse: data,
		},
	}

	for _, tc := range repoStruct {

		mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
		defer mt.Close()

		mt.Run(tc.Name, func(t *mtest.T) {
			if tc.ExpectFindOneErrorReponse != nil {
				t.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
					Message: tc.DataFindOneError.Error(),
				}))
			} else {
				find := mtest.CreateCursorResponse(1, "dbName."+domain.CollectionName, mtest.FirstBatch, bson.D{
					primitive.E{Key: "Id", Value: data.Id},
					primitive.E{Key: "Name", Value: data.Name},
					primitive.E{Key: "Author", Value: data.Author},
				})
				t.AddMockResponses(find)
			}

			if tc.ExpectErrorReponse != nil {
				t.AddMockResponses(mtest.CreateCommandErrorResponse(mtest.CommandError{
					Message: tc.DataError.Error(),
				}))
			} else {
				deleteOne := mtest.CreateSuccessResponse(
					primitive.E{Key: "Id", Value: data.Id},
					primitive.E{Key: "Name", Value: data.Name},
					primitive.E{Key: "Author", Value: data.Author},
				)
				t.AddMockResponses(deleteOne)
			}

			databaseManager := &database.Manager{
				Database: t.DB,
			}
			repo := NewUserMongoRepository(databaseManager, logger)
			err := repo.Delete(context.TODO(), tc.Data)

			if err != nil {

				if tc.ExpectFindOneErrorReponse != nil {
					assert.Contains(t, err.Error(), tc.ExpectFindOneErrorReponse.Error(), tc.Message)
					return
				}
				if tc.ExpectFindOneErrorReponse != nil {
					assert.Contains(t, err.Error(), tc.ExpectErrorReponse.Error(), tc.Message)
					return
				}

				return
			}
		})
	}
}
