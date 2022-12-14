package domain_user_handler

import (
	"time"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	repository "github.com/felixa1996/go_next_be/app/domain/user/repository"
	usecase "github.com/felixa1996/go_next_be/app/domain/user/usecase"
	"github.com/felixa1996/go_next_be/app/infra/database"
)

type UserHandler struct {
	usecase    domain.UserUsecaseContract
	logger     *zap.Logger
	validate   *validator.Validate
	translator ut.Translator
}

func NewUserHandler(usecase domain.UserUsecaseContract, logger *zap.Logger, validate *validator.Validate, translator ut.Translator) UserHandler {
	return UserHandler{
		usecase:    usecase,
		logger:     logger,
		validate:   validate,
		translator: translator,
	}
}

func RegisterUserHandler(db database.Manager, logger *zap.Logger, validate *validator.Validate, translator ut.Translator, contextTimeout time.Duration, group *echo.Group) {
	// init handler
	repo := repository.NewUserMongoRepository(&db, logger)
	usecase := usecase.NewUserUsecase(repo, logger, contextTimeout)
	handler := NewUserHandler(usecase, logger, validate, translator)

	group.GET("", handler.FindPagination)
	group.POST("", handler.Create)
	group.GET("/:id", handler.FindOne)
	group.PATCH("/:id", handler.Update)
	group.DELETE("/:id", handler.Delete)
}
