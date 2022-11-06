package domain_user_usecase

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
	mocks "github.com/felixa1996/go_next_be/mocks/app/domain/user"
)

type testDeleteStruct struct {
	Name               string
	Message            string
	Data               dto.UserDtoDeleteInput
	DataError          error
	DeleteResponse     string
	ExpectErrorReponse error
}

func TestUserDelete(t *testing.T) {
	_, logger := setupTestEnv(t)

	const GENERATED_ID = "31313131-3131-4131-b131-313131313131"

	data := GENERATED_ID

	usecaseStruct := []testDeleteStruct{
		{
			Name:    "Failed",
			Message: "should failed",
			Data: dto.UserDtoDeleteInput{
				Id: GENERATED_ID,
			},
			DeleteResponse:     data,
			DataError:          errors.New("failed to delete user"),
			ExpectErrorReponse: errors.New("failed to delete user"),
		},
		{
			Name:    "Success",
			Message: "should success",
			Data: dto.UserDtoDeleteInput{
				Id: GENERATED_ID,
			},
			DeleteResponse: data,
		},
	}

	for _, tc := range usecaseStruct {
		setupUUID()

		mockRepo := new(mocks.UserRepository)
		if tc.ExpectErrorReponse != nil {
			mockRepo.On("Delete", context.TODO(), tc.DeleteResponse).Return(tc.DataError)
		} else {
			mockRepo.On("Delete", context.TODO(), tc.DeleteResponse).Return(nil)
		}
		usecase := NewUserUsecase(mockRepo, logger, time.Second*2)

		t.Run(tc.Name, func(t *testing.T) {
			err := usecase.Delete(context.TODO(), tc.Data)
			if err != nil {
				assert.Contains(t, err.Error(), tc.ExpectErrorReponse.Error(), tc.Message)
				return
			}
		})
	}
}
