package middlewares

import (
	"context"
	"erply/controllers"
	"erply/infra/database"
	"erply/service"
	"github.com/go-redis/redis/v8"
	"log"
)

func (r *register) resolver() *controllers.Controller {
	// all the external devices
	dh := database.NewRedisHandler(r.redisClient)

	// all the service (usecase/business logic) layer
	cs := service.NewCustomerService(dh)

	// pass to controller (API receiver)
	controller := controllers.NewController(cs)
	return controller
}

type RegisterHandler interface {
	resolver() *controllers.Controller
}

type register struct {
	redisClient *redis.Client
}

func NewRegister(db *redis.Client) RegisterHandler {
	return &register{
		redisClient: db,
	}
}

// setup
// 1. It establishes connection with external devices. Here also gets apikey from .env file to establish connections
// 2. NewRegister() stores all the created connection/client to newRegister
// 3. resolver() sets up all the connections to one controller that contains everything we need
func setup() *controllers.Controller {

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
	// it is does DI (dependency injection)
	controller := newRegister.resolver()

	return controller
}
