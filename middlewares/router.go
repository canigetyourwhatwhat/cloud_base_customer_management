package middlewares

import (
	_ "erply/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func NewRouter() *gin.Engine {

	// get the controller that has everything already set up
	newCon := setup()

	// resolving CORS
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:9000"},
		AllowMethods:     []string{"CREAT", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Authentication
	r.POST("/auth", newCon.Login)

	// customers
	customer := r.Group("/customer")
	customer.POST("create", newCon.CreateCustomer)
	customer.GET(":customerID", newCon.GetCustomerByCustomerID)
	customer.PUT("update", newCon.UpdateCustomer)
	customer.DELETE("delete", newCon.DeleteCustomer)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
