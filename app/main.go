package main

import (
	"github.com/aws/aws-sdk-go/service/sqs"

	"go.uber.org/zap"

	"github.com/felixa1996/go_next_be/app/common"
	. "github.com/felixa1996/go_next_be/app/common"
	"github.com/felixa1996/go_next_be/app/config"
	"github.com/felixa1996/go_next_be/app/infra/database"
	"github.com/felixa1996/go_next_be/app/infra/iam"
	"github.com/felixa1996/go_next_be/app/infra/message"
)

func main() {
	// init config
	config := config.LoadConfigFromEnv()

	// init logging
	logger := InitLogging(config)
	// defer logger.Sync()

	// init database
	dbManager := database.NewDatabaseManager(logger, config.MongoUri, config.MongoDB)

	keycloakIam := iam.NewKeycloakIAM(config)

	// init sqs
	sqsSession, _ := message.NewSqs(logger)
	sqsSvc := sqs.New(sqsSession)

	// init app
	InitApp(config, dbManager, sqsSvc, logger, keycloakIam)

	err := common.Application.Echo.Start(":" + config.Port)
	if err != nil {
		logger.Fatal("Failed to start app", zap.Error(err))
		panic(err)
	}
}
