package main

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/felixa1996/go_next_be/app/common"
	. "github.com/felixa1996/go_next_be/app/common"
	"github.com/felixa1996/go_next_be/app/infra/database"
)

func main() {
	// init config
	ReadConfigFromEnv()

	// init logging
	logger := InitLogging()
	// defer logger.Sync()

	// init database
	dbManager := database.NewDatabaseManager(logger, viper.GetString("MONGODB_URI"), viper.GetString("MONGODB_DB"))

	// init app
	InitApp(dbManager, logger)

	err := common.Application.Echo.Start(":" + viper.GetString("PORT"))
	if err != nil {
		logger.Fatal("Failed to start app",
			zap.String("app", viper.GetString("APP_NAME")),
			zap.Error(err),
		)
		panic(err)
	}
}
