package middlewares

import (
	"erply/controllers"
	_ "erply/docs"
	"erply/service"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"time"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:9000"},
		AllowMethods:     []string{"CREAT", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	newService := service.NewService()
	newCon := controllers.NewController(newService)

	// Authentication
	r.POST("/auth", newCon.Login)

	// customers
	customer := r.Group("/customer")
	customer.POST("create", newCon.CreateCustomer)
	customer.GET(":customerID", newCon.GetCustomerByCustomerID)
	//customer.GET("list", newCon.ListCustomers)
	customer.PUT("update", newCon.UpdateCustomer)
	customer.DELETE("delete", newCon.DeleteCustomer)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
