package common

import (
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.elastic.co/apm/module/apmechov4"
	"go.uber.org/zap"

	"github.com/felixa1996/go_next_be/app/infra/database"
	"github.com/felixa1996/go_next_be/app/infra/iam"
	custom_middleware "github.com/felixa1996/go_next_be/app/infra/middleware"
	"github.com/felixa1996/go_next_be/app/infra/validator"
	_ "github.com/felixa1996/go_next_be/docs"

	domain_company_handler "github.com/felixa1996/go_next_be/app/domain/company/handler"
	domain_user_handler "github.com/felixa1996/go_next_be/app/domain/user/handler"
)

var (
	Application  *App
	appSingleton sync.Once
)

type App struct {
	Echo      *echo.Echo
	dbManager database.Manager
	sqs       *sqs.SQS
	logger    *zap.Logger
	validator.Validator
	keycloakIAM iam.KeycloakIAM
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
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @BasePath /
// @schemes http
func InitApp(dbManager database.Manager, sqs *sqs.SQS, logger *zap.Logger, keycloakIam iam.KeycloakIAM) *App {

	e := echo.New()
	e.HideBanner = true

	e.Use(apmechov4.Middleware())
	e.Use(middleware.CORS())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	appSingleton.Do(func() {
		Application = &App{
			Echo:        e,
			dbManager:   dbManager,
			sqs:         sqs,
			logger:      logger,
			Validator:   *validator.InitValidator(),
			keycloakIAM: keycloakIam,
		}
		Application.RegisterHandlers()
		Application.RegisterEventHandlers()
	})

	return Application
}

// Register REST handler
func (app *App) RegisterHandlers() {
	contextTimeout := time.Duration(viper.GetInt("TIMEOUT")) * time.Second

	user := app.Echo.Group("/v1/user", custom_middleware.KeycloakValidateJwt(app.keycloakIAM))
	domain_user_handler.RegisterUserHandler(app.dbManager, app.logger, app.Validate, app.Translator, contextTimeout, user)
}

// Register Event Handler
func (app *App) RegisterEventHandlers() {
	contextTimeout := time.Duration(viper.GetInt("TIMEOUT")) * time.Second

	domain_company_handler.RegisterCompanyEventHandler(app.dbManager, app.sqs, app.logger, app.Validate, app.Translator, contextTimeout)
}
