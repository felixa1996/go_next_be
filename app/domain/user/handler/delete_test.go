package domain_user_handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
	error_wrapper "github.com/felixa1996/go_next_be/app/infra/error"
	"github.com/felixa1996/go_next_be/app/infra/validator"
	mocks "github.com/felixa1996/go_next_be/mocks/app/domain/user"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type testDeleteStruct struct {
	Name               string
	Message            string
	Data               dto.UserDtoDeleteInput
	DataError          error
	BindingError       bool
	DeleteResponse     interface{}
	ExpectErrorReponse error
	ExpectReponseCode  int
}

func TestUserDelete(t *testing.T) {
	_, logger := setupTestEnv(t)

	validate := validator.InitValidator()

	const GENERATED_ID = "31313131-3131-4131-b131-313131313131"

	usecaseStruct := []testDeleteStruct{
		{
			Name:              "Failed process bind",
			Message:           "should failed process bind",
			BindingError:      true,
			ExpectReponseCode: http.StatusBadRequest,
		},
		{
			Name:         "Failed",
			Message:      "should failed to delete",
			BindingError: false,
			Data: dto.UserDtoDeleteInput{
				Id: GENERATED_ID,
			},
			DataError: error_wrapper.NewErrorWrapper(
				http.StatusBadRequest,
				errors.New("Failed to delete user"),
				"Failed to delete user",
			),
			DeleteResponse: nil,
			ExpectErrorReponse: error_wrapper.NewErrorWrapper(
				http.StatusBadRequest,
				errors.New("Failed to delete user"),
				"Failed to delete user",
			),
			ExpectReponseCode: http.StatusBadRequest,
		},
		{
			Name:         "Success",
			Message:      "should success",
			BindingError: false,
			Data: dto.UserDtoDeleteInput{
				Id: GENERATED_ID,
			},
			ExpectReponseCode: http.StatusNoContent,
		},
	}

	for _, tc := range usecaseStruct {
		t.Run(tc.Name, func(t *testing.T) {
			setupUUID()
			var req *http.Request

			e := echo.New()
			req = httptest.NewRequest(http.MethodDelete, "/", nil)
			if tc.BindingError {
				req = httptest.NewRequest(http.MethodDelete, "/", strings.NewReader("wrong binding"))
			}
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)

			if !tc.BindingError {
				c.SetPath("/v1/user/:id")
				c.SetParamNames("id")
				c.SetParamValues(GENERATED_ID)
			}

			mockUsecase := new(mocks.UserUsecaseContract)
			if tc.ExpectErrorReponse != nil {
				mockUsecase.On("Delete", context.TODO(), tc.Data).Return(tc.DataError)
			} else {
				mockUsecase.On("Delete", context.TODO(), tc.Data).Return(nil)
			}
			handler := NewUserHandler(mockUsecase, logger, validate.Validate, validate.Translator)

			if assert.NoError(t, handler.Delete(c)) {
				assert.Equal(t, tc.ExpectReponseCode, rec.Code)
			}
		})
	}
}
