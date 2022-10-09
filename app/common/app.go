package common

import (
	"sync"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.elastic.co/apm/module/apmechov4"
	"go.uber.org/zap"

	domain_user_handler "github.com/felixa1996/go_next_be/app/domain/user/handler"
	"github.com/felixa1996/go_next_be/app/infra/database"
	_ "github.com/felixa1996/go_next_be/docs"
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

// @title Go Next BE API
// @version 1.0
// @description Sample codebase.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http
func InitApp(dbManager database.Manager, logger *zap.Logger) *App {
	e := echo.New()
	e.Use(apmechov4.Middleware())
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	appSingleton.Do(func() {

		Application = &App{
			Echo:      e,
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
	domain_user_handler.RegisterUserHandler(app.dbManager, app.logger, contextTimeout, user)
}
