package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.uber.org/zap"

	"github.com/felixa1996/go_next_be/app/common"
	. "github.com/felixa1996/go_next_be/app/common"
	"github.com/felixa1996/go_next_be/app/config"
	"github.com/felixa1996/go_next_be/app/infra/database"
	"github.com/felixa1996/go_next_be/app/infra/iam"
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
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	// init app
	InitApp(config, dbManager, sess, logger, keycloakIam)

	err := common.Application.Echo.Start(":" + config.Port)
	if err != nil {
		logger.Fatal("Failed to start app", zap.Error(err))
		panic(err)
	}
}
