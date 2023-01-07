package service

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
)

func (s *CustomerService) CreateCustomer(ctx context.Context, customer *customers.Customer) error {
	err := s.dh.InsertCustomer(ctx, customer)
	if err != nil {
		return err
	}
	return nil
}

func (s *CustomerService) GetCustomerByCustomerID(ctx context.Context, customerId string) (*customers.Customer, error) {
	customer, err := s.dh.GetCustomerByCustomerID(ctx, customerId)
	if err != nil {
		return nil, err
	}
	return customer, nil
}
