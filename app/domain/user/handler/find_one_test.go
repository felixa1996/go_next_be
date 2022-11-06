package domain_user_handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
	error_wrapper "github.com/felixa1996/go_next_be/app/infra/error"
	"github.com/felixa1996/go_next_be/app/infra/validator"
	mocks "github.com/felixa1996/go_next_be/mocks/app/domain/user"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type testFindOneStruct struct {
	Name               string
	Message            string
	Data               dto.UserDtoFindOneInput
	DataError          error
	BindingError       bool
	FindOneResponse    domain.User
	ExpectErrorReponse error
	ExpectReponseCode  int
}

func TestUserFindOne(t *testing.T) {
	fake, logger := setupTestEnv(t)

	validate := validator.InitValidator()

	const GENERATED_ID = "31313131-3131-4131-b131-313131313131"

	data := domain.User{
		Id:     GENERATED_ID,
		Name:   fake.Person().Name(),
		Author: fake.Person().FirstName(),
	}

	usecaseStruct := []testFindOneStruct{
		{
			Name:              "Failed process bind",
			Message:           "should failed process bind",
			BindingError:      true,
			ExpectReponseCode: http.StatusBadRequest,
		},
		{
			Name:         "Failed",
			Message:      "should failed to findone",
			BindingError: false,
			Data: dto.UserDtoFindOneInput{
				Id: GENERATED_ID,
			},
			DataError: error_wrapper.NewErrorWrapper(
				http.StatusBadRequest,
				errors.New("Failed to findone user"),
				"Failed to delete user",
			),
			FindOneResponse: data,
			ExpectErrorReponse: error_wrapper.NewErrorWrapper(
				http.StatusBadRequest,
				errors.New("Failed to findone user"),
				"Failed to findone user",
			),
			ExpectReponseCode: http.StatusBadRequest,
		},
		{
			Name:         "Success",
			Message:      "should success",
			BindingError: false,
			Data: dto.UserDtoFindOneInput{
				Id: GENERATED_ID,
			},
			FindOneResponse:   data,
			ExpectReponseCode: http.StatusOK,
		},
	}

	for _, tc := range usecaseStruct {
		t.Run(tc.Name, func(t *testing.T) {
			setupUUID()
			var req *http.Request

			e := echo.New()
			req = httptest.NewRequest(http.MethodGet, "/", nil)
			if tc.BindingError {
				req = httptest.NewRequest(http.MethodGet, "/", strings.NewReader("wrong binding"))
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
				mockUsecase.On("FindOne", context.TODO(), tc.Data).Return(tc.FindOneResponse, tc.DataError)
			} else {
				mockUsecase.On("FindOne", context.TODO(), tc.Data).Return(tc.FindOneResponse, nil)
			}
			handler := NewUserHandler(mockUsecase, logger, validate.Validate, validate.Translator)

			if assert.NoError(t, handler.FindOne(c)) {
				assert.Equal(t, tc.ExpectReponseCode, rec.Code)
			}
		})
	}
}
