package main

import (
	"github.com/spf13/viper"
	"go.uber.org/zap"

	"github.com/felixa1996/go_next_be/app/common"
	. "github.com/felixa1996/go_next_be/app/common"
	"github.com/felixa1996/go_next_be/app/infra/database"
	"github.com/felixa1996/go_next_be/app/infra/iam"
)

func main() {
	// init config
	ReadConfigFromEnv()

	// init logging
	logger := InitLogging()
	// defer logger.Sync()

	// init database
	dbManager := database.NewDatabaseManager(logger, viper.GetString("MONGODB_URI"), viper.GetString("MONGODB_DB"))

	keycloakIam := iam.NewKeycloakIAM()

	// init app
	InitApp(dbManager, logger, keycloakIam)

	err := common.Application.Echo.Start(":" + viper.GetString("PORT"))
	if err != nil {
		logger.Fatal("Failed to start app", zap.Error(err))
		panic(err)
	}
}
