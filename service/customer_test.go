package service

import (
	"context"
	"erply/entity"
	"erply/infra"
	mockinfra "erply/mocks/infra"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/golang/mock/gomock"
	"reflect"
	"testing"
)

type MockDataHandler struct {
	*mockinfra.MockCustomerHandler
}

func setupDataHandlerTest(t *testing.T) (*MockDataHandler, context.Context) {
	ctx := context.Background()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ch := mockinfra.NewMockCustomerHandler(ctrl)

	// Create a DataHandler with filled interface
	dh := &MockDataHandler{
		MockCustomerHandler: ch,
	}
	return dh, ctx
}

func TestCustomerService_CreateCustomer(t *testing.T) {
	customer := customers.Customer{
		CustomerID:  1,
		FirstName:   "Harry",
		LastName:    "Potter",
		CompanyName: "erply",
	}

	dh, ctx := setupDataHandlerTest(t)
	dh.MockCustomerHandler.EXPECT().InsertCustomer(ctx, &customer).Return(nil)

	type fields struct {
		dh infra.DataHandler
	}
	type args struct {
		ctx      context.Context
		customer *customers.Customer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"insert customer", fields{dh}, args{ctx, &customer}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CustomerService{
				dh: tt.fields.dh,
			}
			if err := s.CreateCustomer(tt.args.ctx, tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("CreateCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCustomerService_GetCustomerByCustomerID(t *testing.T) {

	dh, ctx := setupDataHandlerTest(t)
	customer := customers.Customer{CustomerID: 2}
	dh.MockCustomerHandler.EXPECT().GetCustomerByCustomerID(ctx, "1").Return(&customer, nil)
	dh.MockCustomerHandler.EXPECT().GetCustomerByCustomerID(ctx, "2").Return(nil, entity.ErrCustomerNotFound)

	type fields struct {
		dh infra.DataHandler
	}
	type args struct {
		ctx        context.Context
		customerId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customers.Customer
		wantErr bool
	}{
		{"get existing customer", fields{dh}, args{ctx, "1"}, &customer, false},
		{"get not existing customer", fields{dh}, args{ctx, "2"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &CustomerService{
				dh: tt.fields.dh,
			}
			got, err := s.GetCustomerByCustomerID(tt.args.ctx, tt.args.customerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCustomerByCustomerID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCustomerByCustomerID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
