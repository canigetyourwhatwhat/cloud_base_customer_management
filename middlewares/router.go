package middlewares

import (
	"erply/controllers"
	"erply/docs"
	"erply/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	newService := service.NewService()
	newCon := controllers.NewController(newService)

	// Authentication
	r.POST("/auth", newCon.Login)

	// customers
	customer := r.Group("/customer")
	customer.POST("create", newCon.CreateCustomer)
	customer.GET("get", newCon.GetCustomerByCustomerID)
	customer.PUT("update", newCon.UpdateCustomer)
	customer.DELETE("delete", newCon.DeleteCustomer)

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
