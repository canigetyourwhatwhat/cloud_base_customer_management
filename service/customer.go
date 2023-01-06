package service

import (
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/gin-gonic/gin"
)

func (s *CustomerServiceStruct) CreateCustomer(ctx *gin.Context, customer *customers.Customer) error {
	err := s.dh.InsertCustomer(ctx, customer)
	if err != nil {
		return err
	}
	return nil
}

func (s *CustomerServiceStruct) GetCustomerByCustomerID(ctx *gin.Context, customerId string) (*customers.Customer, error) {
	customer, err := s.dh.GetCustomerByCustomerID(ctx, customerId)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
