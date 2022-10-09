package common

import (
	"sync"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	domain_user "github.com/felixa1996/go_next_be/app/domain/user"
	"github.com/felixa1996/go_next_be/app/infra/database"
)

var (
	Application  *App
	appSingleton sync.Once
)

type App struct {
	Echo      *echo.Echo
	dbManager database.Manager
}

func Init(dbManager database.Manager) *App {
	e := echo.New()
	e.Use(middleware.CORS())

	appSingleton.Do(func() {

		Application = &App{
			Echo:      echo.New(),
			dbManager: dbManager,
		}
		Application.RegisterHandlers()
	})

	return Application
}

// Register REST handler
func (app *App) RegisterHandlers() {

	user := app.Echo.Group("/v1/user")
	domain_user.RegisterUserHandler(app.dbManager, 10000000, user)
}
