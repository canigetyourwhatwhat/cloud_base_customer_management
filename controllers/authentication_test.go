package controllers

import (
	"bytes"
	"encoding/json"
	mockinfra "erply/mocks/infra"
	"erply/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockDataHandler struct {
	*mockinfra.MockCustomerHandler
}

func setupServiceTest(t *testing.T) (*gin.Context, service.CustomerServiceHandler) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ch := mockinfra.NewMockCustomerHandler(ctrl)

	// Create a DataHandler with filled interface
	dh := &MockDataHandler{
		MockCustomerHandler: ch,
	}
	cs := service.NewCustomerService(dh)

	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	return ctx, cs
}

func TestController_Login(t *testing.T) {
	ctx, cs := setupServiceTest(t)
	fail := map[string]string{"username": "", "password": "demo1234", "client_code": "114127"}
	success := map[string]string{"username": "potatochips5050@gmail.com", "password": "demo1234", "client_code": "114127"}

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
		{"login fail with wrong credential", fields{nil, nil, cs}, args{ctx}, fail},
		{"login success with correct credential", fields{nil, nil, cs}, args{ctx}, success},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			con := &Controller{
				sessionKey: tt.fields.sessionKey,
				clientCode: tt.fields.clientCode,
				cs:         tt.fields.cs,
			}
			jsonBytes, err := json.Marshal(tt.body)
			if err != nil {
				t.Error(err)
			}
			tt.args.ctx.Request = &http.Request{
				Body: io.NopCloser(bytes.NewReader(jsonBytes)),
			}
			con.Login(tt.args.ctx)
		})
	}
}
