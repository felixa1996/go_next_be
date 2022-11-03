package database

import (
	"context"

	"github.com/felixa1996/go_next_be/app/infra/healthcheck"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Manager struct {
	Database *mongo.Database
}

func NewDatabaseManager(logger *zap.Logger, uri string, databaseName string) Manager {
	healthcheck.SetMongoReadiness(false)

	clientMongo, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = clientMongo.Ping(context.TODO(), nil)
	if err != nil {
		logger.Fatal("Failed to connect database", zap.Error(err))
		panic(err)
	}
	healthcheck.SetMongoReadiness(true)
	logger.Info("Database connected")

	return Manager{
		Database: clientMongo.Database(databaseName),
	}
}
