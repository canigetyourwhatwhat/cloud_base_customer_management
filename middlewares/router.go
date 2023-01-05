package middlewares

import (
	"erply/controllers"
	"erply/docs"
	"erply/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewRouter(db *sqlx.DB) *gin.Engine {
	r := gin.Default()

	newService := service.NewService(db)
	newCon := controllers.NewController(newService)

	// Authentication
	r.POST("/auth", newCon.Login)

	// customers
	customer := r.Group("/customer")
	customer.POST("create", newCon.CreateCustomer) //

	// @BasePath /api/v1

	// PingExample godoc
	// @Summary ping example
	// @Schemes
	// @Description do ping
	// @Tags example
	// @Accept json
	// @Produce json
	// @Success 200 {string} FetchCustomer
	// @Router /customer/fetch [get]
	customer.GET("fetch", newCon.FetchCustomer) // to fetch the data difference of remote and local server
	customer.GET("get", newCon.GetCustomerByCustomerID)
	customer.PUT("update", newCon.UpdateCustomer)
	customer.DELETE("delete", newCon.DeleteCustomer)

	docs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
