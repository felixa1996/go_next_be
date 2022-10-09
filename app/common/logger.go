package common

import (
	"os"

	"github.com/spf13/viper"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

func InitLogging() *zap.Logger {
	var logger *zap.Logger

	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger = zap.New(core, zap.AddCaller())

	logger = logger.With(zap.String("app.name", viper.GetString("APP_NAME")))

	if viper.GetBool(`debug`) {
		logger.Info("Service run on DEBUG mode", zap.String("app", viper.GetString("APP_NAME")))
	}

	return logger
}
