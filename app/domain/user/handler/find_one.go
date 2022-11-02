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

// UserFindOne godoc
// @Summary      Find One user
// @Description  Find One user
// @Security	  JWT
// @ID get-user-by-id
// @Tags         User
// @Produce      json
// @Param        id 	 path string true "User Id"
// @Success      200  {object}  response.JSONSuccessResult{data=domain_user.User,code=int,message=string}
// @Failure      400  {object}  response.JSONBadRequest{code=int,message=string}
// @Failure      500  {object}  response.JSONInternalServerError{code=int,message=string}
// @Router       /v1/user/{id}  [get]
func (h *UserHandler) FindOne(c echo.Context) error {
	var ew error_wrapper.ErrorWrapper

	ctx := c.Request().Context()

	id := c.Param("id")
	dto := dto.UserDtoDeleteInput{
		Id: id,
	}

	err := h.validate.Struct(dto)
	if err != nil {
		localizedErr := validator.TranslateError(err, h.translator)
		h.logger.Error("Failed to validate find one user", zap.Error(err))
		return response.FailResponse(c, http.StatusBadRequest, "Failed to delete user", localizedErr)
	}

	err = h.usecase.Delete(ctx, dto)
	if err != nil && errors.As(err, &ew) {
		h.logger.Error("Failed to find one user", zap.Error(err))
		return response.FailResponse(c, ew.Code, ew.Message, ew.Err.Error())
	}

	return response.SuccessReponse(c, nil)
}
