package main

import (
	"github.com/gin-gonic/gin"
	logger "github.com/sirupsen/logrus"
	"github.com/Blaze2305/url_short/internal/view"
	"github.com/Blaze2305/url_short/internal/pkg/database"
)

func main(){
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		logger.Print("/")
		c.JSON(200, gin.H{
			"Status": "OK",
		})
	})
	r.Run()
}