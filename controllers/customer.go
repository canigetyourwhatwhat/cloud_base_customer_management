package controllers

import (
	"erply/entity"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
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
	report, err := client.CustomerManager.SaveCustomer(ctx, filter)
	if err != nil {
		ctx.JSON(http.StatusNotAcceptable,
			gin.H{
				"error":         "failed to create a customer in remote server",
				"error content": err.Error(),
			})
		return
	}

	filter["customerID"] = strconv.Itoa(report.CustomerID)

	// Save the customer in the local storage
	err = con.service.CreateCustomer(filter)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error":         "Failed to save data in the local storage",
				"error content": err.Error(),
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "created a customer in remote server and local storage"})
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

	// First, get the customer from the local storage
	customerMissing := false
	localCustomer, err := con.service.GetCustomerByCustomerID(filter["customerID"])

	if errors.Is(entity.ErrCustomerNotFound, err) {
		customerMissing = true
	} else if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error":         "failed to get customer from local server",
				"error_content": err,
			})
		return
	} else {
		// If the customer info is obtained more than 1 hour ago,
		// it updates from the remote Erply server
		if isRecentlyUpdated(localCustomer.UpdatedAt) {
			ctx.JSON(http.StatusOK,
				gin.H{
					"message": "customer info from local storage",
					"result":  localCustomer,
				})
			return
		}
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
				"message":       "this customer doesn't exist",
				"error content": err.Error(),
			})
		return
	}

	filter["firstName"] = remoteCustomers[0].FirstName
	filter["lastName"] = remoteCustomers[0].LastName
	filter["companyName"] = remoteCustomers[0].CompanyName
	filter["email"] = remoteCustomers[0].Email

	if customerMissing {
		err = con.service.CreateCustomer(filter) // update the local data
	} else {
		err = con.service.UpdateCustomerByCustomerID(filter) // update the local data
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":         "failed to update/create customer in local storage",
				"error_content": err,
			})
		return
	}
	log.Printf("customer_id: %s is updated\n", filter["customerID"])

	ctx.JSON(http.StatusOK,
		gin.H{
			"message": "customer info from Erply server",
			"result":  remoteCustomers,
		})
}

func (con *Controller) FetchCustomer(ctx *gin.Context) {
	client, ok := validateUser(ctx, con)
	if !ok {
		return
	}

	customers, err := client.CustomerManager.GetCustomers(ctx, nil)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":         "Failed to fetch data from the Erply server",
				"error content": err.Error(),
			})
		return
	}

	for _, customer := range customers {

		// preprocess the data
		id := strconv.Itoa(customer.CustomerID)
		filter := map[string]string{
			"customerID":  id,
			"firstName":   customer.FirstName,
			"lastName":    customer.LastName,
			"companyName": customer.CompanyName,
			"email":       customer.Email,
		}

		// Store customers if local storage doesn't have compared to remote Erply server
		c, err := con.service.GetCustomerByCustomerID(id)
		if !errors.Is(entity.ErrCustomerNotFound, err) && err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{
					"error":         "Failed to get customers from local storage",
					"error_content": err,
				})
		}

		if errors.Is(entity.ErrCustomerNotFound, err) {
			err = con.service.CreateCustomer(filter)
		} else if isRecentlyUpdated(c.UpdatedAt) {
			err = con.service.UpdateCustomerByCustomerID(filter)
		}

		if err != nil {
			ctx.JSON(http.StatusInternalServerError,
				gin.H{
					"error":         "Failed to save customers in local storage",
					"error_content": err,
				})
		}
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "local storage is the latest version"})
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
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":         "Failed to update data in the Erply server",
				"error content": err.Error(),
			})
		return
	}

	err = con.service.UpdateCustomerByCustomerID(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"error":         "Failed to update data in local storage",
				"error_content": err,
			})
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":         "Failed to delete data in the Erply server",
				"error content": err.Error(),
			})
		return
	}

	err = con.service.DeleteCustomerByCustomerID(filter)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError,
			gin.H{
				"error":         "Failed to delete data in local storage",
				"error_content": err,
			})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "data is deleted from local storage and remote Erply server"})
}
