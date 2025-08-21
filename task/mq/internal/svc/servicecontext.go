package svc

import (
	"context"
	"fmt"
	"go-zero-IM/im/ws/websocket"
	"go-zero-IM/pkg/ctxData"
	"go-zero-IM/task/mq/internal/config"
	"log"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ServiceContext struct {
	config.Config
	Rds      *redis.Client
	MongoDB  *mongo.Database
	WsClient websocket.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	rds := initRedis(c.Rds)
	mongodb := initMongoDB(c.Mongo)

	svc := &ServiceContext{
		Config:  c,
		Rds:     rds,
		MongoDB: mongodb,
	}

	token, err := svc.GetSystemToken()
	if err != nil {
		panic(err)
	}

	header := http.Header{}
	header.Set("Authorization", token)
	svc.WsClient = websocket.NewClient(c.Ws.Host, websocket.WithClientHeader(header))
	return svc
}

func initRedis(cfg config.RedisConfig) *redis.Client {
	rds := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
	})

	fmt.Println("Redis", fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), "connect success")
	return rds
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

func (svc *ServiceContext) GetSystemToken() (string, error) {
	record := svc.Rds.Get(context.Background(), ctxData.REDIS_SYSTEM_ROOT_TOKEN)
	if record == nil {
		return "", record.Err()
	}
	return record.String(), nil
}
