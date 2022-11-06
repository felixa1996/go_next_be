package domain_user_handler

import (
	"bytes"
	"context"
	"encoding/json"
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
	"github.com/google/uuid"
	"github.com/jaswdr/faker"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
)

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

type testCreateStruct struct {
	Name               string
	Message            string
	Data               dto.UserDtoCreateInput
	DataError          error
	BindingError       bool
	CreateResponse     interface{}
	ExpectErrorReponse error
	ExpectReponseCode  int
}

func TestUserCreate(t *testing.T) {
	fake, logger := setupTestEnv(t)

	validate := validator.InitValidator()

	const GENERATED_ID = "31313131-3131-4131-b131-313131313131"

	data := domain.User{
		Id:     GENERATED_ID,
		Name:   fake.Person().Name(),
		Author: fake.Person().FirstName(),
	}

	usecaseStruct := []testCreateStruct{
		{
			Name:              "Failed process bind",
			Message:           "should failed process bind",
			BindingError:      true,
			ExpectReponseCode: http.StatusUnprocessableEntity,
		},
		{
			Name:         "Failed process payload",
			Message:      "should failed process payload",
			BindingError: false,
			Data: dto.UserDtoCreateInput{
				Author: data.Author,
			},
			CreateResponse:    data,
			ExpectReponseCode: http.StatusUnprocessableEntity,
		},
		{
			Name:         "Failed",
			Message:      "should failed to create",
			BindingError: false,
			Data: dto.UserDtoCreateInput{
				Name:   data.Name,
				Author: data.Author,
			},
			DataError: error_wrapper.NewErrorWrapper(
				http.StatusUnprocessableEntity,
				errors.New("Failed to create user"),
				"Failed to create user",
			),
			CreateResponse: data,
			ExpectErrorReponse: error_wrapper.NewErrorWrapper(
				http.StatusUnprocessableEntity,
				errors.New("Failed to create user"),
				"Failed to create user",
			),
			ExpectReponseCode: http.StatusUnprocessableEntity,
		},
		{
			Name:         "Success",
			Message:      "should success",
			BindingError: false,
			Data: dto.UserDtoCreateInput{
				Name:   data.Name,
				Author: data.Author,
			},
			CreateResponse:    data,
			ExpectReponseCode: http.StatusCreated,
		},
	}

	for _, tc := range usecaseStruct {
		t.Run(tc.Name, func(t *testing.T) {
			var req *http.Request

			payloadDataJson, _ := json.Marshal(tc.Data)
			payloadDataString := string(payloadDataJson)

			e := echo.New()
			req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payloadDataString))
			if tc.BindingError {
				req = httptest.NewRequest(http.MethodPost, "/", strings.NewReader("wrong binding"))
			}
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			setupUUID()

			mockUsecase := new(mocks.UserUsecaseContract)
			if tc.ExpectErrorReponse != nil {
				mockUsecase.On("Create", context.TODO(), tc.Data).Return(tc.CreateResponse, tc.DataError)
			} else {
				mockUsecase.On("Create", context.TODO(), tc.Data).Return(tc.CreateResponse, nil)
			}
			handler := NewUserHandler(mockUsecase, logger, validate.Validate, validate.Translator)

			if assert.NoError(t, handler.Create(c)) {
				assert.Equal(t, tc.ExpectReponseCode, rec.Code)
			}
		})
	}
}
