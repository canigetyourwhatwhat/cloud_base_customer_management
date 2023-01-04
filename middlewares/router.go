package middlewares

import (
	"erply/controllers"
	"erply/service"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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
	customer.GET("fetch", newCon.FetchCustomer)    // to fetch the data difference of remote and local server
	customer.GET("get", newCon.GetCustomerByCustomerID)
	customer.PUT("update", newCon.UpdateCustomer)
	customer.DELETE("delete", newCon.DeleteCustomer)

	return r
}
