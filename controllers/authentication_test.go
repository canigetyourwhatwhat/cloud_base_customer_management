package controllers

import (
	"erply/service"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestController_Login(t *testing.T) {
	//cs := service.NewCustomerService()
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := &Controller{
				sessionKey: tt.fields.sessionKey,
				clientCode: tt.fields.clientCode,
				cs:         tt.fields.cs,
			}
			con.Login(tt.args.ctx)
		})
	}
}
