package main

import (
	"os"

	"github.com/Blaze2305/url_short/internal/pkg/database"
	"github.com/Blaze2305/url_short/internal/view"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

func registerV1Routes(ctx *gin.RouterGroup, routeProvider view.ProviderMethods) {
	ctx.POST("/urls", routeProvider.CreateURL)
	// ctx.GET("/urls", routeProvider.ListURLs)

}

func registerBaseRoutes(ctx *gin.RouterGroup, routeProvider view.ProviderMethods) {
	ctx.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]string{"Status": "OK"})
	})

	ctx.GET("/:id", routeProvider.Redirect)

}

func main() {
	r := gin.Default()

	r.Use(location.Default())
	baseGroup := r.Group("/")
	v1Group := r.Group("/v1")
	db := database.NewDatabase(os.Getenv("URL_MONGO"))
	v1Provider := view.NewProvider(db)
	baseProvider := view.NewProvider(db)

	registerV1Routes(v1Group, v1Provider)
	registerBaseRoutes(baseGroup, baseProvider)
	r.Run()
}
