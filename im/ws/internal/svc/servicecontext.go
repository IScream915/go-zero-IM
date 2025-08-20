package svc

import (
	"context"
	"go-zero-IM/im/ws/internal/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceContext struct {
	Config  config.Config
	MongoDB *mongo.Database

	//immodels.ChatLogModel
	//mqclient.MsgChatTransferClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	mongodb := initMongoDB(c.Mongo)

	return &ServiceContext{
		Config:  c,
		MongoDB: mongodb,
		//MsgChatTransferClient: mqclient.NewMsgChatTransferClient(c.MsgChatTransfer.Addrs, c.MsgChatTransfer.Topic),
		//ChatLogModel:          immodels.MustChatLogModel(c.Mongo.Url, c.Mongo.Db),
	}
}

func initMongoDB(cfg config.MongoConfig) *mongo.Database {
	log.Printf("Connecting to MongoDB at %s...", cfg.Url)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.Url))
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Ping the database to verify connection
	if err = client.Ping(ctx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	log.Printf("Successfully connected to MongoDB database: %s", cfg.Db)
	return client.Database(cfg.Db)
}
