package database

import (
	"bytes"
	"context"
	"encoding/gob"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/go-redis/redis/v8"
	"log"
	"reflect"
	"strconv"
	"testing"
	"time"
)

var db = ConnectDB()
var ctx = context.Background()

func ConnectDB() *redis.Client {
	// Connect Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		log.Println("failed to connect redis")
		panic(err)
	}

	return redisClient
}

func Test_customerHandler_GetCustomerByCustomerID(t *testing.T) {

	// Insert the data in the database
	customer := customers.Customer{
		CustomerID:  1,
		FirstName:   "Harry",
		LastName:    "Potter",
		CompanyName: "erply",
	}
	var buff bytes.Buffer
	if err := gob.NewEncoder(&buff).Encode(customer); err != nil {
		t.Error(err)
	}
	customerID := strconv.Itoa(customer.CustomerID)
	if err := db.Set(ctx, customerID, buff.Bytes(), 1*time.Hour).Err(); err != nil {
		t.Error(err)
	}

	type fields struct {
		db *redis.Client
	}
	type args struct {
		ctx        context.Context
		customerID string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *customers.Customer
		wantErr bool
	}{
		{"get value when there is data", fields{db}, args{ctx, customerID}, &customer, false},
		{"get value when there isn't data", fields{db}, args{ctx, "2"}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := &CustomerHandler{
				db: tt.fields.db,
			}
			got, err := ch.GetCustomerByCustomerID(ctx, tt.args.customerID)
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

func Test_customerHandler_InsertCustomer(t *testing.T) {
	customer := customers.Customer{
		CustomerID:  1,
		FirstName:   "Harry",
		LastName:    "Potter",
		CompanyName: "erply",
	}

	type fields struct {
		db *redis.Client
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
		{"check if the value is inserted", fields{db}, args{ctx, &customer}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ch := &CustomerHandler{
				db: tt.fields.db,
			}
			if err := ch.InsertCustomer(tt.args.ctx, tt.args.customer); (err != nil) != tt.wantErr {
				t.Errorf("InsertCustomer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
