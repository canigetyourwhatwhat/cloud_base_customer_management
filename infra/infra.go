package infra

import (
	"context"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
)

// DataHandler
// interface for whole external devices
type DataHandler interface {
	CustomerHandler
}

// CustomerHandler
// interface for handling Customer information with Redis
type CustomerHandler interface {
	InsertCustomer(ctx context.Context, customer *customers.Customer) error
	GetCustomerByCustomerID(ctx context.Context, customerID string) (*customers.Customer, error)
}
