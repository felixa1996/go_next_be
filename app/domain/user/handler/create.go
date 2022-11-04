package domain_user_handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
	error_wrapper "github.com/felixa1996/go_next_be/app/infra/error"
	"github.com/felixa1996/go_next_be/app/infra/response"
	"github.com/felixa1996/go_next_be/app/infra/validator"
)

// UserCreate godoc
// @Summary      Create user
// @Description  Create user
// @Security	  JWT
// @Tags         User
// @Produce      json
// @Param        user body domain_user_dto.UserDtoCreateInput true "User Data"
// @Success      201  {object}  response.JSONSuccessResult{data=domain_user.User,code=int,message=string}
// @Failure      400  {object}  response.JSONBadRequest{code=int,message=string}
// @Failure      422  {object}  response.JSONUnprocessableEntity{code=int,message=string}
// @Failure      500  {object}  response.JSONInternalServerError{code=int,message=string}
// @Router       /v1/user [post]
func (h *UserHandler) Create(c echo.Context) error {
	var ew error_wrapper.ErrorWrapper

	ctx := c.Request().Context()

	var dto dto.UserDtoCreateInput

	err := c.Bind(&dto)
	if err != nil {
		h.logger.Error("Failed to process payload create user", zap.Error(err))
		return response.FailResponse(c, http.StatusUnprocessableEntity, "Failed to process payload", err.Error())
	}

	err = h.validate.Struct(dto)
	if err != nil {
		localizedErr := validator.TranslateError(err, h.translator)
		h.logger.Error("Failed to validate create user", zap.Error(err))
		return response.FailResponse(c, http.StatusUnprocessableEntity, "Failed to process entity", localizedErr)
	}

	res, err := h.usecase.Create(ctx, dto)
	if err != nil && errors.As(err, &ew) {
		h.logger.Error("Failed to create user", zap.Error(err))
		return response.FailResponse(c, ew.Code, ew.Message, ew.Err.Error())
	}

	return response.SuccessCreatedReponse(c, res)
}
