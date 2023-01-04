package entity

import "errors"

// Used in the backend
var (
	ErrCustomerNotFound = errors.New("ths customers doesn't exist")
)

// Used to return front-end as error code
const (
	Err_Customer_Not_Found  = "CUSTOMER_NOT_FOUND"
	Err_Input_Invalud       = "INPUT_INVALID"
	Err_Validation_Failed   = "Err_Validation_Failed"
	Err_Parsing_JSON_Failed = "Err_Parsing_JSON_Failed"
	Err_Customers_Not_Found = "Err_Customers_Not_Found"
)
