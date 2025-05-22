package persistence

import (
	"context"
	"log"
	"rest_services_with_http_net/configs"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	once        sync.Once
)

func InitMongoClient() {
	once.Do(func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		clientOptions := options.Client().
			ApplyURI(configs.AppConfig.MongoURI).
			SetConnectTimeout(time.Duration(configs.AppConfig.ConnectionTimeout) * time.Millisecond).
			SetServerSelectionTimeout(time.Duration(configs.AppConfig.ServerSelectionTimeout) * time.Millisecond).
			SetSocketTimeout(time.Duration(configs.AppConfig.SocketTimeout) * time.Millisecond).
			SetMaxPoolSize(uint64(configs.AppConfig.MaxPoolSize)).
			SetMinPoolSize(uint64(configs.AppConfig.MinPoolSize)).
			SetMaxConnIdleTime(time.Duration(configs.AppConfig.MaxConnIdleTime) * time.Second)
		client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
			log.Fatal("Failed to connect to MongoDB:", err)
		}
		mongoClient = client
		log.Println("Connected to MongoDB")
	})
}

func GetCollection(name string) *mongo.Collection {
	if mongoClient == nil {
		log.Fatal("MongoDB client is not initialized")
	}
	db := mongoClient.Database(configs.AppConfig.DBName)
	return db.Collection(name)
}
