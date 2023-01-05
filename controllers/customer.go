package controllers

import (
	"erply/entity"
	"fmt"
	"github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (con *Controller) CreateCustomer(ctx *gin.Context) {

	// Parse the body
	var body entity.Customer
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error_code": entity.Err_Parsing_JSON_Failed,
			})
		return
	}

	client, ok := validateUser(ctx, con)
	if !ok {
		return
	}

	// Save the customer
	filter := map[string]string{
		"firstName":   body.FirstName,
		"lastName":    body.LastName,
		"companyName": body.CompanyName,
		"email":       body.Email,
	}

	// Save the customer in the remote Erply server
	_, err := client.CustomerManager.SaveCustomer(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable,
			gin.H{
				"error":         "failed to create a customer in remote server",
				"error content": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created a customer"})
}

func (con *Controller) GetCustomerByCustomerID(ctx *gin.Context) {

	// Parse the body
	var body entity.Customer
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error_code": entity.Err_Parsing_JSON_Failed,
			})
		return
	}

	client, ok := validateUser(ctx, con)
	if !ok {
		return
	}
	filter := map[string]string{
		"customerID": body.CustomerID,
	}

	// First, get the customer from the cache
	localCustomer, err := con.service.GetCustomerByCustomerID(ctx, filter["customerID"])
	if err != nil {
		fmt.Println(err)
	}
	if err == nil {
		ctx.JSON(http.StatusOK,
			gin.H{
				"message": "customer info from cache",
				"result":  localCustomer,
			})
		return
	}

	// Get customer from remote Erply server to update/create customer info in local storage
	remoteCustomers, err := client.CustomerManager.GetCustomers(ctx, filter)
	if len(remoteCustomers) == 0 {
		ctx.JSON(http.StatusOK,
			gin.H{"message": "this customer doesn't exist"})
		return
	}
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable,
			gin.H{
				"error":         "Failed to get customer data from Erply server",
				"error content": err.Error(),
			})
		return
	}

	err = con.service.CreateCustomer(ctx, &remoteCustomers[0])
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"error":         "Failed to save customer data in redis",
				"error content": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK,
		gin.H{
			"message": "customer info from Erply server",
			"result":  remoteCustomers,
		})
}

func handleCustomerError(err error) error {
	if erplyError, ok := err.(*common.ErplyError); ok {
		switch erplyError.Code {
		case 1011:
			return entity.ErrCustomerNotFound
		default:
			return err
		}
	} else {
		return err
	}
}

func (con *Controller) UpdateCustomer(ctx *gin.Context) {
	// Parse the body
	var body entity.Customer
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error_code": entity.Err_Parsing_JSON_Failed,
			})
		return
	}

	client, ok := validateUser(ctx, con)
	if !ok {
		return
	}

	filter := map[string]string{
		"customerID":  body.CustomerID,
		"firstName":   body.FirstName,
		"lastName":    body.LastName,
		"companyName": body.CompanyName,
		"email":       body.Email,
	}

	_, err := client.CustomerManager.SaveCustomer(ctx, filter)
	err = handleCustomerError(err)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":         "Failed to update data in the Erply server",
				"error content": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "data is updated in local storage and remote Erply server"})
}

func (con *Controller) DeleteCustomer(ctx *gin.Context) {

	// Parse the body
	var body entity.Customer
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error_code": entity.Err_Parsing_JSON_Failed,
			})
		return
	}

	client, ok := validateUser(ctx, con)
	if !ok {
		return
	}

	filter := map[string]string{
		"customerID":  body.CustomerID,
		"firstName":   body.FirstName,
		"lastName":    body.LastName,
		"companyName": body.CompanyName,
		"email":       body.Email,
	}

	err := client.CustomerManager.DeleteCustomer(ctx, filter)
	if err != nil {
		err = handleCustomerError(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":         "Failed to delete data in the Erply server",
				"error content": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "data is deleted from local storage and remote Erply server"})
}
