package response

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type JSONSuccessResult struct {
	Code    int         `json:"code" example:"200"`
	Message string      `json:"message" example:"Success"`
	Data    interface{} `json:"data"`
}

type JSONBadRequest struct {
	Code    int         `json:"code" example:"400"`
	Message string      `json:"message" example:"Bad Request"`
	Data    interface{} `json:"data" `
}

type JSONUnprocessableEntity struct {
	Code    int         `json:"code" example:"422"`
	Message string      `json:"message" example:"Unprocessable Entity"`
	Data    interface{} `json:"data" `
}

type JSONInternalServerError struct {
	Code    int         `json:"code" example:"500"`
	Message string      `json:"message" example:"Internal Server Error"`
	Data    interface{} `json:"data" `
}

// nolint
func SuccessReponse(c echo.Context, data interface{}) error {
	c.JSON(http.StatusOK, JSONSuccessResult{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    data,
	})
	return nil
}

// nolint
func SuccessCreatedReponse(c echo.Context, data interface{}) error {
	c.JSON(http.StatusCreated, JSONSuccessResult{
		Code:    http.StatusCreated,
		Message: "Success",
		Data:    data,
	})
	return nil
}

// nolint
func SuccessNoContentReponse(c echo.Context, data interface{}) error {
	c.JSON(http.StatusNoContent, JSONSuccessResult{
		Code:    http.StatusNoContent,
		Message: "Success",
		Data:    data,
	})
	return nil
}

// nolint
func FailResponse(c echo.Context, code int, message string, data interface{}) error {
	if code == http.StatusInternalServerError {
		c.JSON(code, JSONInternalServerError{
			Code:    code,
			Data:    data,
			Message: message,
		})
		return nil
	}

	c.JSON(code, JSONBadRequest{
		Code:    code,
		Data:    data,
		Message: message,
	})
	return nil
}
