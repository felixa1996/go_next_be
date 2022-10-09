package database

import (
	"context"

	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Manager struct {
	Client *mongo.Client
}

func (manager *Manager) InitDB(uri string) error {
	clientMongo, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = clientMongo.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	log.Info("Database connected")

	manager.Client = clientMongo
	return nil
}
