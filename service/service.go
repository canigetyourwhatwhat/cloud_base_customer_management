package service

import (
	"context"
	"erply/infra"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
)

//go:generate mockgen -source=service/service.go --destination=mocks/service/serviceMockInterface.go service CustomerServiceHandler
type CustomerServiceHandler interface {
	CreateCustomer(ctx context.Context, customer *customers.Customer) error
	GetCustomerByCustomerID(ctx context.Context, customerId string) (*customers.Customer, error)
}

type CustomerService struct {
	dh infra.DataHandler
}

func NewCustomerService(dh infra.DataHandler) CustomerServiceHandler {
	return &CustomerService{
		dh: dh,
	}
}
