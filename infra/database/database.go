package database

import (
	"bytes"
	"encoding/gob"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func InsertCustomer(ctx *gin.Context, redisDB *redis.Client, customer *customers.Customer) error {

	var buff bytes.Buffer
	customerIdStr := strconv.Itoa(customer.CustomerID)

	if err := gob.NewEncoder(&buff).Encode(customer); err != nil {
		return err
	}
	if err := redisDB.Set(ctx, customerIdStr, buff.Bytes(), 1*time.Hour).Err(); err != nil {
		return err
	}

	return nil
}

func GetCustomerByCustomerID(ctx *gin.Context, redisDB *redis.Client, customerID string) (*customers.Customer, error) {
	cmd := redisDB.Get(ctx, customerID)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(cmdb)
	var customer customers.Customer
	if err = gob.NewDecoder(reader).Decode(&customer); err != nil {
		return nil, err
	}

	return &customer, nil
}
