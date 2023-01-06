package database

import (
	"bytes"
	"encoding/gob"
	"erply/infra"
	"github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

func NewRedisHandler(db *redis.Client) infra.DataHandler {
	ch := NewCustomerHandler(db)

	return &redisHandler{ch}
}

type redisHandler struct {
	infra.CustomerHandler
}

type customerHandler struct {
	db *redis.Client
}

func NewCustomerHandler(db *redis.Client) infra.CustomerHandler {
	return &customerHandler{db}
}

func (ch *customerHandler) InsertCustomer(ctx *gin.Context, customer *customers.Customer) error {

	// First, it converts the customer data into byte code using gob pacakge
	var buff bytes.Buffer
	if err := gob.NewEncoder(&buff).Encode(customer); err != nil {
		return err
	}

	// Next, it stores the encoded customer data as a value and customerID as a key into Redis
	customerIdStr := strconv.Itoa(customer.CustomerID)
	if err := ch.db.Set(ctx, customerIdStr, buff.Bytes(), 1*time.Hour).Err(); err != nil {
		return err
	}

	return nil
}

func (ch *customerHandler) GetCustomerByCustomerID(ctx *gin.Context, customerID string) (*customers.Customer, error) {

	// Get the encoded customer data from Redis
	cmd := ch.db.Get(ctx, customerID)

	// Convert it as Bytes because it was stored as bytes
	cmdb, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}

	// Creates bytes reader and decode the encoded customer data into entity.Customer type
	reader := bytes.NewReader(cmdb)
	var customer customers.Customer
	if err = gob.NewDecoder(reader).Decode(&customer); err != nil {
		return nil, err
	}

	return &customer, nil
}
