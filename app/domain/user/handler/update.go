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
// @ID get-user-by-id
// @Param        id 	 path string true "User Id"
// @Param        user body domain_user_dto.UserDtoUpdateInput true "User Data"
// @Success      200  {object}  response.JSONSuccessResult{data=domain_user.User,code=int,message=string}
// @Failure      400  {object}  response.JSONBadRequest{code=int,message=string}
// @Failure      422  {object}  response.JSONUnprocessableEntity{code=int,message=string}
// @Failure      500  {object}  response.JSONInternalServerError{code=int,message=string}
// @Router       /v1/user/{id} [patch]
func (h *UserHandler) Update(c echo.Context) error {
	var ew error_wrapper.ErrorWrapper

	ctx := c.Request().Context()

	id := c.Param("id")
	dtoParams := dto.UserDtoUpdateParamInput{
		Id: id,
	}

	err := h.validate.Struct(dtoParams)
	if err != nil {
		localizedErr := validator.TranslateError(err, h.translator)
		h.logger.Error("Failed to validate find one user", zap.Error(err))
		return response.FailResponse(c, http.StatusBadRequest, "Failed to delete user", localizedErr)
	}

	var dto dto.UserDtoUpdateInput

	err = c.Bind(&dto)
	if err != nil {
		h.logger.Error("Failed to process payload update user", zap.Error(err))
		return response.FailResponse(c, http.StatusUnprocessableEntity, "Failed to process payload", err.Error())
	}

	err = h.validate.Struct(dto)
	if err != nil {
		localizedErr := validator.TranslateError(err, h.translator)
		h.logger.Error("Failed to validate update user", zap.Error(err))
		return response.FailResponse(c, http.StatusUnprocessableEntity, "Failed to process entity", localizedErr)
	}

	res, err := h.usecase.Update(ctx, dtoParams, dto)
	if err != nil && errors.As(err, &ew) {
		h.logger.Error("Failed to update user", zap.Error(err))
		return response.FailResponse(c, ew.Code, ew.Message, ew.Err.Error())
	}

	return response.SuccessReponse(c, res)
}
