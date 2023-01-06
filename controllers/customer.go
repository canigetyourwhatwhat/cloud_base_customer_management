package controllers

import (
	"erply/entity"
	"fmt"
	_ "github.com/erply/api-go-wrapper/pkg/api/customers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// CreateCustomer
//	@Summary		For user to create new customer data
//	@Description	It creates new customer data in remote Erply server
//	@Param			entity.Customer	body	entity.Customer	true	"Customer data"
//	@Accept			json
//	@Router			/customer/create [post]
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

	client, err := validateUser(con)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error":   entity.Err_Validation_Failed,
				"message": "Please login again"})
		return
	}

	// Save the customer
	filter := map[string]string{
		"customerID":  body.CustomerID,
		"firstName":   body.FirstName,
		"lastName":    body.LastName,
		"companyName": body.CompanyName,
		"email":       body.Email,
	}

	// Save the customer in the remote Erply server
	_, err = client.CustomerManager.SaveCustomer(ctx, filter)
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

// GetCustomerByCustomerID
//	@Summary		For user to get existing customer data
//	@Description	It gets customer existing customer data from cache if there is data, if not, from the remote Erply server
//	@Param			customerID	path	string	true	"customer id"
//	@Accept			json
//	@Router			/customer/{customerID} [get]
func (con *Controller) GetCustomerByCustomerID(ctx *gin.Context) {

	customerID := ctx.Param("customerID")

	client, err := validateUser(con)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error":   entity.Err_Validation_Failed,
				"message": "Please login again"})
		return
	}

	filter := map[string]string{
		"customerID": customerID,
	}

	// First, get the customer from the cache
	localCustomer, err := con.cs.GetCustomerByCustomerID(ctx, filter["customerID"])
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
				"error":         "Failed to get customer data from remote Erply server",
				"error content": err.Error(),
			})
		return
	}

	err = con.cs.CreateCustomer(ctx, &remoteCustomers[0])
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
			"message": "customer info from remote Erply server",
			"result":  remoteCustomers,
		})
}

// UpdateCustomer
//	@Summary		For user to update existing customer data
//	@Description	It updates existing customer data in remote Erply server, and it doesn't store this change in local storage.
//	@Param			entity.Customer	body	entity.Customer	true	"Customer data"
//	@Accept			json
//	@Router			/customer/update [put]
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

	client, err := validateUser(con)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error":   entity.Err_Validation_Failed,
				"message": "Please login again"})
		return
	}

	filter := map[string]string{
		"customerID":  body.CustomerID,
		"firstName":   body.FirstName,
		"lastName":    body.LastName,
		"companyName": body.CompanyName,
		"email":       body.Email,
	}

	_, err = client.CustomerManager.SaveCustomer(ctx, filter)
	err = handleCustomerError(err)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":         "Failed to update data in the remote Erply server",
				"error content": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "data is updated for remote Erply server"})
}

// DeleteCustomer
//	@Summary		For user to delete existing customer data
//	@Description	It deletes existing customer data in remote Erply server, and it doesn't store this change in local storage.
//	@Param			entity.Customer	body	entity.Customer	true	"Customer data"
//	@Accept			json
//	@Router			/customer/delete [delete]
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

	client, err := validateUser(con)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error":   entity.Err_Validation_Failed,
				"message": "Please login again"})
		return
	}

	filter := map[string]string{
		"customerID":  body.CustomerID,
		"firstName":   body.FirstName,
		"lastName":    body.LastName,
		"companyName": body.CompanyName,
		"email":       body.Email,
	}

	err = client.CustomerManager.DeleteCustomer(ctx, filter)
	if err != nil {
		err = handleCustomerError(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":         "Failed to delete data in the remote Erply server",
				"error content": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "data is deleted from remote Erply server"})
}
