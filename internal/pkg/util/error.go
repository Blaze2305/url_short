package util

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

// HTTPError - send http request error
func HTTPError(c *gin.Context, errorCode int, err error) {
	log.Errorf(err.Error())
	c.JSON(
		errorCode,
		gin.H{"error": err.Error()},
	)
}
