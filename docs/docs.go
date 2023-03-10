// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Daichi Ando",
            "url": "https://github.com/canigetyourwhatwhat/cloud_base_customer_management/blob/main/README.md",
            "email": "daichiando98@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth": {
            "post": {
                "description": "It is required to login to call other endpoints",
                "consumes": [
                    "application/json"
                ],
                "summary": "For user to login",
                "parameters": [
                    {
                        "description": "Data to login",
                        "name": "entity.AuthenticationInfo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.AuthenticationInfo"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/customer/create": {
            "post": {
                "description": "It creates new customer data in remote Erply server",
                "consumes": [
                    "application/json"
                ],
                "summary": "For user to create new customer data",
                "parameters": [
                    {
                        "description": "Customer data",
                        "name": "entity.Customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Customer"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/customer/delete": {
            "delete": {
                "description": "It deletes existing customer data in remote Erply server, and it doesn't store this change in local storage.",
                "consumes": [
                    "application/json"
                ],
                "summary": "For user to delete existing customer data",
                "parameters": [
                    {
                        "description": "Customer data",
                        "name": "entity.Customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Customer"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/customer/update": {
            "put": {
                "description": "It updates existing customer data in remote Erply server, and it doesn't store this change in local storage.",
                "consumes": [
                    "application/json"
                ],
                "summary": "For user to update existing customer data",
                "parameters": [
                    {
                        "description": "Customer data",
                        "name": "entity.Customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/entity.Customer"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/customer/{customerID}": {
            "get": {
                "description": "It gets customer existing customer data from cache if there is data, if not, from the remote Erply server",
                "consumes": [
                    "application/json"
                ],
                "summary": "For user to get existing customer data",
                "parameters": [
                    {
                        "type": "string",
                        "description": "customer id",
                        "name": "customerID",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "entity.AuthenticationInfo": {
            "type": "object",
            "properties": {
                "client_code": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "entity.Customer": {
            "type": "object",
            "properties": {
                "companyName": {
                    "type": "string"
                },
                "customerID": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:9000",
	BasePath:         "",
	Schemes:          []string{"http"},
	Title:            "Erply cache server",
	Description:      "It reads and writes customer data using Erply API. It sues cache with Redis.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
