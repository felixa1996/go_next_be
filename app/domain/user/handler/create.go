package domain_user_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	dto "github.com/felixa1996/go_next_be/app/domain/user/dto"
	"github.com/felixa1996/go_next_be/app/infra/validator"
)

// UserCreate godoc
// @Summary      Create user
// @Description  Create user
// @Tags         User
// @Produce      json
// @Param        user body domain_user_dto.UserDtoCreateInput true "User Data"
// @Success      200  {array}  domain_user.User
// @Router       /v1/user [post]
func (h *UserHandler) Create(c echo.Context) error {

	ctx := c.Request().Context()

	var userDto dto.UserDtoCreateInput

	err := c.Bind(&userDto)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	err = h.validate.Struct(userDto)
	if err != nil {
		localizedErr := validator.TranslateError(err, h.translator)
		h.logger.Error("Failed to validate create user", zap.Error(err))
		return c.JSON(http.StatusUnprocessableEntity, localizedErr)
	}

	res, err := h.usecase.Create(ctx, userDto)
	if err != nil {
		h.logger.Error("Failed to create user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}
