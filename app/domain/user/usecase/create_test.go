package domain_user_usecase

import (
	"bytes"
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
	mocks "github.com/felixa1996/go_next_be/mocks/app/domain/user"
)

type testStruct struct {
	Name                 string
	Message              string
	Data                 dto.UserDtoCreateInput
	DataError            error
	CreateResponse       interface{}
	ExpectSuccessReponse domain.User
	ExpectErrorReponse   error
}

func setupTestEnv(t *testing.T) (faker.Faker, *zap.Logger) {
	reader := bytes.NewReader([]byte("1111111111111111"))
	uuid.SetRand(reader)
	uuid.SetClockSequence(1)

	fake := faker.New()
	logger := zaptest.NewLogger(t)
	return fake, logger
}

func TestUserCreate(t *testing.T) {
	t.Parallel()
	fake, logger := setupTestEnv(t)

	const GENERATED_ID = "31313131-3131-4131-b131-313131313131"

	data := domain.User{
		Id:     GENERATED_ID,
		Name:   fake.Person().Name(),
		Author: fake.Person().FirstName(),
	}

	usecaseStruct := []testStruct{
		{
			Name:    "Success",
			Message: "should success",
			Data: dto.UserDtoCreateInput{
				Name:   data.Name,
				Author: data.Author,
			},
			CreateResponse:       data,
			ExpectSuccessReponse: data,
		},
		// {
		// 	Name:    "Failed",
		// 	Message: "should failed",
		// 	Data: dto.UserDtoCreateInput{
		// 		Name:   data.Name,
		// 		Author: data.Author,
		// 	},
		// 	CreateResponse:     nil,
		// 	DataError:          errors.New("failed to create user"),
		// 	ExpectErrorReponse: errors.New("failed to create user"),
		// },
	}

	for _, tc := range usecaseStruct {
		mockRepo := new(mocks.UserRepository)
		if tc.ExpectErrorReponse != nil {
			mockRepo.On("Create", context.TODO(), tc.ExpectErrorReponse).Return(nil, tc.DataError)
		} else {
			mockRepo.On("Create", context.TODO(), tc.CreateResponse).Return(tc.CreateResponse, nil)
		}
		usecase := NewUserUsecase(mockRepo, logger, time.Second*2)

		t.Run(tc.Name, func(t *testing.T) {
			res, err := usecase.Create(context.TODO(), tc.Data)
			if err != nil {
				assert.Contains(t, err.Error(), tc.ExpectErrorReponse.Error(), tc.Message)
				return
			}

			assert.Equal(t, tc.ExpectSuccessReponse, res, tc.Message)
		})
	}

	// t.Run("Success", func(y *testing.T) {

	// 	mockRepo := new(mocks.UserRepository)
	// 	mockRepo.On("Create", context.TODO(), user).Return(user, nil)

	// 	a := NewUserUsecase(mockRepo, logger, time.Second*2)
	// 	res, err := a.Create(context.TODO(), dto.UserDtoCreateInput{
	// 		Name:   user.Name,
	// 		Author: user.Author,
	// 	})
	// 	if err != nil {
	// 		// assert.Contains(t, err.Error(), tc.ExpectErrorReponse.Error(), "Should error")
	// 	}

	// 	assert.Equal(t, user, res, "Should success")
	// })
}
