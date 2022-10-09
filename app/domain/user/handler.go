package domain_user

import (
	"log"
	"net/http"
	"time"

	"github.com/felixa1996/go_next_be/app/infra/database"
	"github.com/labstack/echo/v4"
)

func RegisterUserHandler(db database.Manager, contextTimeout time.Duration, group *echo.Group) {
	// init used repository and usecase
	repo := NewUserMongoRepository(&db)
	// todo need change
	usecase := NewUserUsecase(repo, contextTimeout)

	group.GET("", func(c echo.Context) error {
		ctx := c.Request().Context()
		res, err := usecase.FindPagination(ctx)
		if err != nil {
			log.Fatal("Failed to fetch user", err)
			return c.JSON(http.StatusInternalServerError, err)
		}
		return c.JSON(http.StatusOK, res)
	})
}
