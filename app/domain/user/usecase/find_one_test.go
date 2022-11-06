package domain_user_usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
	mocks "github.com/felixa1996/go_next_be/mocks/app/domain/user"
)

type testFindOneStruct struct {
	Name                 string
	Message              string
	Data                 dto.UserDtoFindOneInput
	DataError            error
	FindOneResponse      interface{}
	ExpectSuccessReponse domain.User
	ExpectErrorReponse   error
}

func TestUserFindOne(t *testing.T) {
	fake, logger := setupTestEnv(t)

	const GENERATED_ID = "31313131-3131-4131-b131-313131313131"

	data := domain.User{
		Id:     GENERATED_ID,
		Name:   fake.Person().Name(),
		Author: fake.Person().FirstName(),
	}

	usecaseStruct := []testFindOneStruct{
		{
			Name:    "Failed",
			Message: "should failed",
			Data: dto.UserDtoFindOneInput{
				Id: GENERATED_ID,
			},
			FindOneResponse:    data,
			DataError:          errors.New("failed to findone user"),
			ExpectErrorReponse: errors.New("failed to findone user"),
		},
		{
			Name:    "Success",
			Message: "should success",
			Data: dto.UserDtoFindOneInput{
				Id: GENERATED_ID,
			},
			FindOneResponse:      data,
			ExpectSuccessReponse: data,
		},
	}

	for _, tc := range usecaseStruct {
		setupUUID()

		mockRepo := new(mocks.UserRepository)
		if tc.ExpectErrorReponse != nil {
			mockRepo.On("FindOne", context.TODO(), tc.Data.Id).Return(tc.FindOneResponse, tc.DataError)
		} else {
			mockRepo.On("FindOne", context.TODO(), tc.Data.Id).Return(tc.FindOneResponse, nil)
		}
		usecase := NewUserUsecase(mockRepo, logger, time.Second*2)

		t.Run(tc.Name, func(t *testing.T) {
			res, err := usecase.FindOne(context.TODO(), tc.Data)
			if err != nil {
				assert.Contains(t, err.Error(), tc.ExpectErrorReponse.Error(), tc.Message)
				return
			}

			assert.Equal(t, tc.ExpectSuccessReponse, res, tc.Message)
		})
	}
}
