package common

import (
	"sync"
	"time"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"go.elastic.co/apm/module/apmechov4"
	"go.uber.org/zap"

	"github.com/felixa1996/go_next_be/app/config"
	"github.com/felixa1996/go_next_be/app/infra/database"
	"github.com/felixa1996/go_next_be/app/infra/healthcheck"
	"github.com/felixa1996/go_next_be/app/infra/iam"
	"github.com/felixa1996/go_next_be/app/infra/message"
	custom_middleware "github.com/felixa1996/go_next_be/app/infra/middleware"
	"github.com/felixa1996/go_next_be/app/infra/tracer"
	"github.com/felixa1996/go_next_be/app/infra/uploader"
	"github.com/felixa1996/go_next_be/app/infra/validator"
	_ "github.com/felixa1996/go_next_be/docs"

	domain_company_handler "github.com/felixa1996/go_next_be/app/domain/company/handler"
	domain_media_handler "github.com/felixa1996/go_next_be/app/domain/media/handler"
	domain_user_handler "github.com/felixa1996/go_next_be/app/domain/user/handler"
)

var (
	Application *App
	once        sync.Once
)

type App struct {
	Config       *config.Config
	Echo         *echo.Echo
	dbManager    database.Manager
	sqsWrapper   message.SqsWrapper
	minioWrapper uploader.MinioWrapper
	logger       *zap.Logger
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

// @host localhost:3004
// @securityDefinitions.apiKey JWT
// @in header
// @name Authorization
// @BasePath /
// @schemes http
func InitApp(config config.Config, dbManager database.Manager, sess *session.Session, minioWrapper uploader.MinioWrapper, logger *zap.Logger, keycloakIam iam.KeycloakIAM) *App {

	// sqs
	sqsWrapper := message.NewSqsWrapper(logger, sess)
	e := echo.New()
	e.HideBanner = true

	switch traceType := config.TraceType; traceType {
	case "newrelic":
		newRelicApp := tracer.NewRelicTracer(config)
		e.Use(nrecho.Middleware(newRelicApp))
	case "elk":
		e.Use(apmechov4.Middleware())
	}

	// Enable metrics middleware
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
	e.Use(apmechov4.Middleware())
	e.Use(middleware.CORS())

	e.GET("/docs/*", echoSwagger.WrapHandler)

	once.Do(func() {
		Application = &App{
			Config:       &config,
			Echo:         e,
			dbManager:    dbManager,
			sqsWrapper:   sqsWrapper,
			minioWrapper: minioWrapper,
			logger:       logger,
			Validator:    *validator.InitValidator(),
			keycloakIAM:  keycloakIam,
		}
		Application.RegisterHandlers()
		// Register as go routine so it can block the main thread
		go Application.RegisterEventHandlers()
	})

	return Application
}

// Register REST handler
func (app *App) RegisterHandlers() {
	contextTimeout := time.Duration(app.Config.Timeout) * time.Second

	probe := app.Echo.Group("/probes")
	healthcheck.RegisterHealthCheckHandler(app.logger, probe)

	user := app.Echo.Group("/v1/user", custom_middleware.KeycloakValidateJwt(app.keycloakIAM))
	domain_user_handler.RegisterUserHandler(app.dbManager, app.logger, app.Validate, app.Translator, contextTimeout, user)

	media := app.Echo.Group("/v1/media", custom_middleware.KeycloakValidateJwt(app.keycloakIAM))
	domain_media_handler.RegisterMediaHandler(app.dbManager, app.minioWrapper, app.logger, app.Validate, app.Translator, contextTimeout, media)

	app.logger.Info("Successfully register REST handler")
}

// Register Event Handler
func (app *App) RegisterEventHandlers() {
	contextTimeout := time.Duration(app.Config.Timeout) * time.Second

	domain_company_handler.RegisterCompanyEventHandler(app.Config, app.dbManager, app.sqsWrapper, app.logger, app.Validate, app.Translator, contextTimeout)

	app.logger.Info("Successfully register Event handler")
}
