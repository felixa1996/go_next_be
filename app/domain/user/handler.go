package domain_user

import (
	"net/http"
	"time"

	"github.com/felixa1996/go_next_be/app/infra/database"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func RegisterUserHandler(db database.Manager, logger *zap.Logger, contextTimeout time.Duration, group *echo.Group) {
	// init used repository and usecase
	repo := NewUserMongoRepository(&db, logger)
	usecase := NewUserUsecase(repo, logger, contextTimeout)

	group.GET("", func(c echo.Context) error {
		ctx := c.Request().Context()
		res, err := usecase.FindPagination(ctx)
		if err != nil {
			logger.Fatal("Failed to fetch user", zap.Error(err))
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, res)
	})
}
