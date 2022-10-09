package domain_user_handler

import (
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	domain "github.com/felixa1996/go_next_be/app/domain/user"
	repository "github.com/felixa1996/go_next_be/app/domain/user/repository"
	usecase "github.com/felixa1996/go_next_be/app/domain/user/usecase"
	"github.com/felixa1996/go_next_be/app/infra/database"
)

type UserHandler struct {
	usecase domain.UserUsecaseContract
	logger  *zap.Logger
}

func RegisterUserHandler(db database.Manager, logger *zap.Logger, contextTimeout time.Duration, group *echo.Group) {
	// init handler
	repo := repository.NewUserMongoRepository(&db, logger)
	usecase := usecase.NewUserUsecase(repo, logger, contextTimeout)
	handler := &UserHandler{
		usecase: usecase,
		logger:  logger,
	}

	group.GET("", handler.FindPagination)
}
