package connector

import (
	"context"
	"movie-service/pkg/logger"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Options struct {
	Uri string
}

func InitMongo(opts Options) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(opts.Uri))
	if err != nil {
		logger.Fatalf("failed to connect mongo: %s\n", err.Error())
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logger.Fatalf("failed to ping mongo: %s\n", err.Error())
	}

	logger.Info("Connected to MongoDB!")

	return client
}
