package database

import (
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	connectTimeOut         = 10
	errorConnectionMessage = "error connecting to MongoDB: %v"
)

var (
	mongoClient *mongo.Client
	mongoOnce   sync.Once
)

type MongoConnectionPort interface {
	Connection() *mongo.Client
	Close() error
	CheckConnection() int
}

type MongoConnection struct{}

func (*MongoConnection) Connection() *mongo.Database {
	dbName := os.Getenv("MONGO_DATABASE")
	if mongoClient == nil {
		connection, err := createConnection()
		if err != nil {
			panic(fmt.Errorf("error connecting to MongoDB: %v", err))
		}
		return connection.Database(dbName)
	}
	return mongoClient.Database(dbName)
}

func createConnection() (*mongo.Client, error) {
	mongoOnce.Do(func() {
		dbPort := os.Getenv("MONGO_PORT")
		dbHost := os.Getenv("MONGO_HOST")
		connectionUri := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
		mongoClient = connect(connectionUri)
	})
	return mongoClient, nil
}

func connect(uri string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithTimeout(context.Background(), connectTimeOut*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(fmt.Errorf(errorConnectionMessage, err))
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(fmt.Errorf(errorConnectionMessage, err))
	}

	fmt.Println("🍃 Connected to MongoDB! 🍃")

	return client
}

func (*MongoConnection) CheckConnection() int {
	err := mongoClient.Ping(context.Background(), nil)
	if err != nil {
		return 0
	}
	return 1
}

func (*MongoConnection) Close() error {
	if mongoClient != nil {
		return mongoClient.Disconnect(context.Background())
	}
	return nil
}
