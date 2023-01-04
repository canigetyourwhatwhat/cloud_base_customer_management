package middlewares

import (
	"fmt"
	erply "github.com/erply/api-go-wrapper/pkg/api"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func Prep() *gin.Context {

	// todo create the DI structs here

	// Get the variables from .env file to authenticate the users
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")
	clientCode := os.Getenv("CLIENT_CODE")

	client, err := erply.NewClientFromCredentials(username, password, clientCode, nil)
	//client, err := api.Client{AuthProvider: username}
	if err != nil {
		fmt.Println("Authentication is failed")
		fmt.Println(err)
	}

	ctx := gin.Context{Keys: map[string]interface{}{"client": client}}

	return &ctx
}
