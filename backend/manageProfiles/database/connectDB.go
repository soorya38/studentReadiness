package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectToDb() (context.Context, *mongo.Client, context.CancelFunc, error) {
	mongoURI := "mongodb+srv://test:test@cluster0.208wr.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	if err != nil {
		cancel() // Cancel context if error occurs
		return nil, nil, nil, fmt.Errorf("failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(ctx, nil); err != nil {
		cancel() // Cancel context if error occurs
		return nil, nil, nil, fmt.Errorf("failed to ping MongoDB: %v", err)
	}

	return ctx, client, cancel, nil
}
