package entity

import "errors"

// Used in the backend
var (
	ErrCustomerNotFound           = errors.New("ths customer doesn't exist in the Erply server")
	ErrLoginInfoMissing           = errors.New("session key or/and Client code is missing")
	ErrFailedEstablishErplyClient = errors.New("failed to establish client")
)

// Used to return front-end as error code
const (
	Err_Customer_Not_Found  = "CUSTOMER_NOT_FOUND"
	Err_Input_Invalud       = "INPUT_INVALID"
	Err_Validation_Failed   = "Err_Validation_Failed"
	Err_Parsing_JSON_Failed = "Err_Parsing_JSON_Failed"
)
