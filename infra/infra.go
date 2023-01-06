package infra

import (
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/gin-gonic/gin"
)

// DataHandler
// interface for whole external devices
type DataHandler interface {
	CustomerHandler
}

// CustomerHandler
// interface for handling Customer information with Redis
type CustomerHandler interface {
	InsertCustomer(ctx *gin.Context, customer *customers.Customer) error
	GetCustomerByCustomerID(ctx *gin.Context, customerID string) (*customers.Customer, error)
}
