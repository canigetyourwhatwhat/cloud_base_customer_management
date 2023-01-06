package service

import (
	"erply/infra/database"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api/customers"

	"github.com/gin-gonic/gin"
)

func (s *Service) CreateCustomer(ctx *gin.Context, customer *customers.Customer) error {
	err := database.InsertCustomer(ctx, s.redisDB, customer)
	if err != nil {
		fmt.Println("hit")
		return err
	}
	fmt.Println("good")
	return nil
}

func (s *Service) GetCustomerByCustomerID(ctx *gin.Context, customerId string) (*customers.Customer, error) {
	customer, err := database.GetCustomerByCustomerID(ctx, s.redisDB, customerId)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
