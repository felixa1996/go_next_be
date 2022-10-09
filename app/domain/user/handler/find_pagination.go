package domain_user_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// UserFindPagination godoc
// @Summary      Find user pagination
// @Description  Find user pagination
// @Tags         User
// @Produce      json
// @Success      200  {array}  domain_user.User
// @Router       /v1/user [get]
func (h *UserHandler) FindPagination(c echo.Context) error {
	ctx := c.Request().Context()
	res, err := h.usecase.FindPagination(ctx)
	if err != nil {
		h.logger.Error("Failed to fetch user", zap.Error(err))
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, res)
}
