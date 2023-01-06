package controllers

import (
	"erply/service"
	"github.com/gin-gonic/gin"
)

type Controller struct {
	sessionKey *string
	clientCode *string
	cs         service.CustomerServiceInterface
}

type ControllerHandler interface {
	CreateCustomer(ctx *gin.Context)
	GetCustomerByCustomerID(ctx *gin.Context)
	UpdateCustomer(ctx *gin.Context)
	DeleteCustomer(ctx *gin.Context)
}

func NewController(cs service.CustomerServiceInterface) *Controller {
	return &Controller{
		cs: cs,
	}
}
