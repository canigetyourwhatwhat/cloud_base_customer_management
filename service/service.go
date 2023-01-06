package service

import (
	"erply/infra"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/gin-gonic/gin"
)

type CustomerServiceInterface interface {
	CreateCustomer(ctx *gin.Context, customer *customers.Customer) error
	GetCustomerByCustomerID(ctx *gin.Context, customerId string) (*customers.Customer, error)
}

type CustomerServiceStruct struct {
	dh infra.DataHandler
}

func NewCustomerService(dh infra.DataHandler) CustomerServiceInterface {
	return &CustomerServiceStruct{
		dh: dh,
	}
}
