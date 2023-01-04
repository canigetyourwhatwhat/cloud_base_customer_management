package controllers

import (
	erply "github.com/erply/api-go-wrapper/pkg/api"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const TimeSpanByHours = 1.0

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

// It returns false if the target is not updated less than 1 hour
func isRecentlyUpdated(target time.Time) bool {
	return time.Now().Sub(target).Seconds() < TimeSpanByHours
}

// Removes all the nil values
//func removeNils(m map[string]string) {
//	val := reflect.ValueOf(m)
//	for _, e := range val.MapKeys() {
//		v := val.MapIndex(e)
//		if !v.IsValid() {
//			delete(m, e.String())
//		}
//	}
//}
