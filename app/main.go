package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"

	"go.uber.org/zap"

	"github.com/felixa1996/go_next_be/app/common"
	. "github.com/felixa1996/go_next_be/app/common"
	"github.com/felixa1996/go_next_be/app/config"
	"github.com/felixa1996/go_next_be/app/infra/database"
	"github.com/felixa1996/go_next_be/app/infra/healthcheck"
	"github.com/felixa1996/go_next_be/app/infra/iam"
	"github.com/felixa1996/go_next_be/app/infra/uploader"
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
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	if err != nil {
		healthcheck.SetAwsSessionReadiness(false)
		logger.Fatal("Failed to initialize aws session", zap.Error(err))
		panic(err)
	}
	healthcheck.SetAwsSessionReadiness(true)

	// init minio
	minio := uploader.NewMinioWrapper(config, logger)

	// init app
	InitApp(config, dbManager, sess, minio, logger, keycloakIam)

	go func() {
		err = common.Application.Echo.Start(":" + config.Port)
		if err != nil && err != http.ErrServerClosed {
			logger.Fatal("F ailed to start app", zap.Error(err))
		}
		println(err)
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := common.Application.Echo.Shutdown(ctx); err != nil {
		logger.Fatal("Something wrong when shutdown server", zap.Error(err))
	}
}
