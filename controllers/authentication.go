package controllers

import (
	"erply/entity"
	"github.com/erply/api-go-wrapper/pkg/api/auth"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (con *Controller) Login(ctx *gin.Context) {
	var body entity.AuthenticationInfo
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error_code": entity.Err_Input_Invalud,
			})
		return
	}

	var httpCli http.Client
	sessionKey, err := auth.VerifyUser(body.Username, body.Password, body.ClientCode, &httpCli)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error_code": entity.Err_Validation_Failed,
			})
		return
	}

	con.sessionKey = &sessionKey
	con.clientCode = &body.ClientCode

	ctx.JSON(http.StatusOK, gin.H{
		"message": "authentication was successful",
	})
}
