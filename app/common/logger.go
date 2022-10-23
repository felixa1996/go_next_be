package common

import (
	"os"

	"github.com/felixa1996/go_next_be/app/config"
	"go.elastic.co/ecszap"
	"go.uber.org/zap"
)

func InitLogging(config config.Config) *zap.Logger {
	var logger *zap.Logger

	encoderConfig := ecszap.NewDefaultEncoderConfig()
	core := ecszap.NewCore(encoderConfig, os.Stdout, zap.DebugLevel)
	logger = zap.New(core, zap.AddCaller())

	logger = logger.With(zap.String("app.name", config.AppName))

	if config.Debug {
		logger.Info("Service run on DEBUG mode", zap.String("app", config.AppName))
	}

	return logger
}
