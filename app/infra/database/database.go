package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

type Manager struct {
	Client   *mongo.Client
	Database *mongo.Database
}

func NewDatabaseManager(logger *zap.Logger, uri string, databasName string) Manager {
	clientMongo, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = clientMongo.Ping(context.TODO(), nil)
	if err != nil {
		logger.Fatal("Failed to connect database", zap.String("err", err.Error()))
		panic(err)
	}
	logger.Info("Database connected")

	return Manager{
		Client:   clientMongo,
		Database: clientMongo.Database(databasName),
	}
}

func (manager *Manager) InitDB(uri string, databaseName string) {

}
