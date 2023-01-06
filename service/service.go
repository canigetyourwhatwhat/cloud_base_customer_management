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
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Println("failed to connect redis")
		panic(err)
		return nil
	}
	return &Service{redisDB: client}
}
