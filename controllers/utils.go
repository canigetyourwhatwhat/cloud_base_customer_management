package controllers

import (
	"erply/entity"
	erply "github.com/erply/api-go-wrapper/pkg/api"
	"github.com/erply/api-go-wrapper/pkg/api/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func validateUser(ctx *gin.Context, con *Controller) (*erply.Client, bool) {

	if con.sessionKey == nil || con.clientCode == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			gin.H{
				"error":   "Session key or/and Client code is missing",
				"message": "Please login again"})
		return nil, false
	}
	httpCli := http.Client{}
	client, err := erply.NewClient(*con.sessionKey, *con.clientCode, &httpCli)
	if err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"error":   "Failed to establish client",
				"message": "Please login again"})
		return nil, false
	}
	return client, true
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
