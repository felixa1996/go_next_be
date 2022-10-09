package common

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func InitLogging() *zap.Logger {
	var logger *zap.Logger

	if viper.GetBool(`debug`) {
		logger, _ = zap.NewDevelopment()
		logger.Info("Service run on DEBUG mode", zap.String("app", viper.GetString("APP_NAME")))
	} else {
		logger, _ = zap.NewProduction()
	}
	return logger
}
