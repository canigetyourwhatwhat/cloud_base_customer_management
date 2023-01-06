package middlewares

import (
	"context"
	"erply/controllers"
	"erply/infra/database"
	"erply/service"
	"github.com/go-redis/redis/v8"
	"log"
)

func (r *registerStruct) NewResolver() *controllers.Controller {
	// all the external devices
	dh := database.NewRedisHandler(r.redisClient)

	// all the service (usecase/business logic) layer
	cs := service.NewCustomerService(dh)

	// pass to controller (API receiver)
	controller := controllers.NewController(cs)
	return controller
}

type RegisterInterface interface {
	NewResolver() *controllers.Controller
}

type registerStruct struct {
	redisClient *redis.Client
}

func NewRegister(db *redis.Client) RegisterInterface {
	return &registerStruct{
		redisClient: db,
	}
}

//
//func BuildComponents() *controllers.Controller {
//	newService := service.NewService()
//	newCon := controllers.NewController(newService)
//	return newCon
//}

func RegisterExternalDevices() *controllers.Controller {

	// Connect Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.Println("failed to connect redis")
		panic(err)
	}

	// store all the registered components in newRegister
	newRegister := NewRegister(redisClient)

	// store all the registered components to the controller struct to be able to use all of them
	// it is dependency injection
	controller := newRegister.NewResolver()

	return controller

}
