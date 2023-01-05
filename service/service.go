package service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

type Service struct {
	redisDB *redis.Client
}

func NewService() *Service {
	//client := redis.NewClient(&redis.Options{
	//	Addr:        "localhost:6379",
	//	DB:          0,
	//	DialTimeout: 100 * time.Millisecond,
	//	ReadTimeout: 100 * time.Millisecond,
	//})

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Println("failed to connect redis")
		panic(err)
		return nil
	}
	return &Service{redisDB: client}
}
