package view

import (
	"log"
	"time"

	"github.com/Blaze2305/url_short/internal/pkg/constants"
	"github.com/Blaze2305/url_short/internal/pkg/model"
	"github.com/Blaze2305/url_short/internal/pkg/util"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

// CreateURL - create a shortened url and return it
func (p Provider) CreateURL(c *gin.Context) {
	input := model.Shorten{}

	c.BindJSON(&input)

	token := util.NewToken(5)
	if input.Forward[:4] != "http" {
		input.Forward = "https://" + input.Forward
	}
	log.Print(token)
	input.Token = token
	input.Created = time.Now().String()

	_, err := p.db.CreateURL(input)
	if err != nil {
		util.HTTPError(c, constants.BadRequestCode, err)
		return
	}

	c.JSON(200, map[string]string{"url": location.Get(c).String() + "/" + token})
}

// Redirect the new url to the original url
func (p Provider) Redirect(c *gin.Context) {
	token := c.Param("id")

	urlObj, err := p.db.GetURL(token)
	if err != nil {
		util.HTTPError(c, constants.BadRequestCode, err)
		return
	}
	log.Printf("%#v", *urlObj)
	log.Print(c.Request.Host)
	log.Print(c.Request.RequestURI)
	// c.JSON(200, map[string]string{"Status": "OK", "Url": *urlObj})
	// c.Redirect(http.StatusMovedPermanently, "http://google.co.in")
	// c.Writer.Header().Set("Location", "http://www.google.com")
	c.Redirect(308, *urlObj)
	c.Abort()
}
