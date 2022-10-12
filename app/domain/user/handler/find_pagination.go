package domain_user_handler

import (
	"errors"

	error_wrapper "github.com/felixa1996/go_next_be/app/infra/error"
	"github.com/felixa1996/go_next_be/app/infra/response"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// UserFindPagination godoc
// @Summary      Find user pagination
// @Description  Find user pagination
// @Security	  JWT
// @Tags         User
// @Produce      json
// @Success      200  {object} response.JSONSuccessResult{data=[]domain_user.User,code=int,message=string}
// @Failure      500  {object} response.JSONInternalServerError{code=int,message=string}
// @Router       /v1/user [get]
func (h *UserHandler) FindPagination(c echo.Context) error {
	var ew error_wrapper.ErrorWrapper
	ctx := c.Request().Context()
	res, err := h.usecase.FindPagination(ctx)

	if err != nil && errors.As(err, &ew) {
		h.logger.Error("Failed to fetch user", zap.Error(err))
		return response.FailResponse(c, ew.Code, ew.Message, ew.Err.Error())
	}
	return response.SuccessReponse(c, res)
}
