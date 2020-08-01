package main

import (
	"os"

	"github.com/Blaze2305/url_short/internal/pkg/database"
	"github.com/Blaze2305/url_short/internal/view"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

// Handle the routes for the application
func registerRoutes(ctx *gin.RouterGroup, routeProvider view.ProviderMethods) {

	// Base route to check for server status
	ctx.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]string{"Status": "OK"})
	})

	// Redirect route to redirect to forwarding url from token
	ctx.GET("/:id", routeProvider.Redirect)

	// URL related route : create , delete urls
	ctx.POST("/urls", routeProvider.CreateURL)
	ctx.DELETE("/:id/url", routeProvider.DeleteURL)

	// User related routes : create , get , delete users
	ctx.POST("/users", routeProvider.CreateUser)
	ctx.GET("/:id/user", routeProvider.GetUser)
	ctx.DELETE("/:id/user", routeProvider.DeleteUser)
}

func main() {
	r := gin.Default()

	// ***Uncomment when going to prod***
	// gin.SetMode(gin.ReleaseMode)

	// location used to get the hostname of the current machine the app is hosted on
	// so that I dont need to hardcode it if I ever change hosting platforms
	r.Use(location.Default())

	// Groups together all the routes under one base route group
	baseGroup := r.Group("/")

	// Create a new database pointint to the given mongo atlas cluster
	db := database.NewDatabase(os.Getenv("URL_MONGO"))

	// Createa a new provider to interact with db
	baseProvider := view.NewProvider(db)

	registerRoutes(baseGroup, baseProvider)
	r.Run()
}
