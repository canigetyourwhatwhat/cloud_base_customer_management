package controllers

import (
	"bytes"
	"encoding/json"
	"erply/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"net/url"
	"testing"
)

func TestController_CreateCustomer(t *testing.T) {
	ctx, cs := setupServiceTest(t)
	credential := map[string]string{"username": "", "password": "demo1234", "client_code": "114127"}
	input := map[string]string{
		"firstName":   "harry1",
		"lastName":    "potter2",
		"companyName": "Test",
		"email":       "test@gmail.com",
	}

	type fields struct {
		sessionKey *string
		clientCode *string
		cs         service.CustomerServiceHandler
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		body   map[string]string
	}{
		{"success with correct credential and valid input", fields{nil, nil, cs}, args{ctx}, input},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := &Controller{
				sessionKey: tt.fields.sessionKey,
				clientCode: tt.fields.clientCode,
				cs:         tt.fields.cs,
			}

			// login
			jsonBytes, err := json.Marshal(credential)
			if err != nil {
				t.Error(err)
			}
			tt.args.ctx.Request = &http.Request{
				Body: io.NopCloser(bytes.NewReader(jsonBytes)),
			}
			con.Login(tt.args.ctx)

			// make the API call
			jsonBytes, err = json.Marshal(tt.body)
			if err != nil {
				t.Error(err)
			}
			tt.args.ctx.Request = &http.Request{
				Body: io.NopCloser(bytes.NewReader(jsonBytes)),
			}
			con.CreateCustomer(tt.args.ctx)
		})
	}
}

func TestController_DeleteCustomer(t *testing.T) {
	ctx, cs := setupServiceTest(t)
	credential := map[string]string{"username": "", "password": "demo1234", "client_code": "114127"}
	input := map[string]string{
		"firstName":   "harry1",
		"lastName":    "potter2",
		"companyName": "Test",
		"email":       "test@gmail.com",
	}

	type fields struct {
		sessionKey *string
		clientCode *string
		cs         service.CustomerServiceHandler
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		body   map[string]string
	}{
		{"success with correct credential and valid input", fields{nil, nil, cs}, args{ctx}, input},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := &Controller{
				sessionKey: tt.fields.sessionKey,
				clientCode: tt.fields.clientCode,
				cs:         tt.fields.cs,
			}

			// login
			jsonBytes, err := json.Marshal(credential)
			if err != nil {
				t.Error(err)
			}
			tt.args.ctx.Request = &http.Request{
				Body: io.NopCloser(bytes.NewReader(jsonBytes)),
			}
			con.Login(tt.args.ctx)

			// make the API call
			jsonBytes, err = json.Marshal(tt.body)
			if err != nil {
				t.Error(err)
			}
			tt.args.ctx.Request = &http.Request{
				Body: io.NopCloser(bytes.NewReader(jsonBytes)),
			}
			con.DeleteCustomer(tt.args.ctx)
		})
	}
}

func TestController_GetCustomerByCustomerID(t *testing.T) {
	ctx, cs := setupServiceTest(t)
	credential := map[string]string{"username": "", "password": "demo1234", "client_code": "114127"}

	type fields struct {
		sessionKey *string
		clientCode *string
		cs         service.CustomerServiceHandler
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"success with correct credential and valid input", fields{nil, nil, cs}, args{ctx}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := &Controller{
				sessionKey: tt.fields.sessionKey,
				clientCode: tt.fields.clientCode,
				cs:         tt.fields.cs,
			}
			// login
			jsonBytes, err := json.Marshal(credential)
			if err != nil {
				t.Error(err)
			}
			tt.args.ctx.Request = &http.Request{
				Body: io.NopCloser(bytes.NewReader(jsonBytes)),
			}
			con.Login(tt.args.ctx)

			// make the API call
			tt.args.ctx.Request = nil
			params := []gin.Param{
				{
					Key:   "customerID",
					Value: "1",
				},
			}
			tt.args.ctx.Params = params
			u := url.Values{}
			u.Add("customerID", "1")
			tt.args.ctx.Request = &http.Request{
				URL: &url.URL{Host: "localhost:9000", Path: "customer"},
			}
			con.GetCustomerByCustomerID(tt.args.ctx)
		})
	}
}

func TestController_UpdateCustomer(t *testing.T) {
	ctx, cs := setupServiceTest(t)
	credential := map[string]string{"username": "", "password": "demo1234", "client_code": "114127"}
	input := map[string]string{
		"firstName":   "harry1",
		"lastName":    "potter2",
		"companyName": "Test",
		"email":       "test@gmail.com",
	}

	type fields struct {
		sessionKey *string
		clientCode *string
		cs         service.CustomerServiceHandler
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		body   map[string]string
	}{
		{"success with correct credential and valid input", fields{nil, nil, cs}, args{ctx}, input},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := &Controller{
				sessionKey: tt.fields.sessionKey,
				clientCode: tt.fields.clientCode,
				cs:         tt.fields.cs,
			}

			// login
			jsonBytes, err := json.Marshal(credential)
			if err != nil {
				t.Error(err)
			}
			tt.args.ctx.Request = &http.Request{
				Body: io.NopCloser(bytes.NewReader(jsonBytes)),
			}
			con.Login(tt.args.ctx)

			// make the API call
			jsonBytes, err = json.Marshal(tt.body)
			if err != nil {
				t.Error(err)
			}
			tt.args.ctx.Request = &http.Request{
				Body: io.NopCloser(bytes.NewReader(jsonBytes)),
			}
			con.UpdateCustomer(tt.args.ctx)
		})
	}
}
