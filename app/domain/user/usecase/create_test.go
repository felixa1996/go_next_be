package domain_user_usecase

import (
	"bytes"
	"context"
	"errors"
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

type testCreateStruct struct {
	Name                 string
	Message              string
	Data                 dto.UserDtoCreateInput
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

func setupUUID() {
	reader := bytes.NewReader([]byte("1111111111111111"))
	uuid.SetRand(reader)
	uuid.SetClockSequence(1)
}

func TestUserCreate(t *testing.T) {
	fake, logger := setupTestEnv(t)

	const GENERATED_ID = "31313131-3131-4131-b131-313131313131"

	data := domain.User{
		Id:     GENERATED_ID,
		Name:   fake.Person().Name(),
		Author: fake.Person().FirstName(),
	}

	usecaseStruct := []testCreateStruct{
		{
			Name:    "Failed",
			Message: "should failed",
			Data: dto.UserDtoCreateInput{
				Name:   data.Name,
				Author: data.Author,
			},
			CreateResponse:     data,
			DataError:          errors.New("failed to create user"),
			ExpectErrorReponse: errors.New("failed to create user"),
		},
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
	}

	for _, tc := range usecaseStruct {
		setupUUID()

		mockRepo := new(mocks.UserRepository)
		if tc.ExpectErrorReponse != nil {
			mockRepo.On("Create", context.TODO(), tc.CreateResponse).Return(tc.CreateResponse, tc.DataError)
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
}
