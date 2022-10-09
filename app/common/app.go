package common

import (
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"go.elastic.co/apm/module/apmechov4"
	"go.uber.org/zap"

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
	logger    *zap.Logger
}

func InitApp(dbManager database.Manager, logger *zap.Logger) *App {
	e := echo.New()
	e.Use(apmechov4.Middleware())
	e.Use(middleware.CORS())

	appSingleton.Do(func() {

		Application = &App{
			Echo:      echo.New(),
			dbManager: dbManager,
			logger:    logger,
		}
		Application.RegisterHandlers()
	})

	return Application
}

// Register REST handler
func (app *App) RegisterHandlers() {
	contextTimeout := time.Duration(viper.GetInt("TIMEOUT")) * time.Second

	user := app.Echo.Group("/v1/user")
	domain_user.RegisterUserHandler(app.dbManager, app.logger, contextTimeout, user)
}
