package main

import (
	"flag"
	"os"

	"github.com/Blaze2305/url_short/internal/pkg/database"
	"github.com/Blaze2305/url_short/internal/view"
	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
)

var ginMode = flag.String("mode", "debug", "set gin mode to run in [debug | test | release] default = debug ")

// Handle the routes for the application
func registerRoutes(ctx *gin.RouterGroup, routeProvider view.ProviderMethods) {

	// Base route to check for server status
	ctx.GET("/", func(c *gin.Context) {
		c.JSON(200, map[string]string{"Status": "OK"})
	})

	// Route to redirect to forwarding url from token
	ctx.GET("/:id", routeProvider.Redirect)

	// URL related route : create , delete urls
	ctx.POST("/urls", routeProvider.CreateURL)
	ctx.DELETE("/:id/url", routeProvider.DeleteURL)

	// User related routes : create , get , delete , edit users
	ctx.POST("/users", routeProvider.CreateUser)
	ctx.GET("/:id/user", routeProvider.GetUser)
	ctx.DELETE("/:id/user", routeProvider.DeleteUser)
	ctx.PUT("/:id/user", routeProvider.UpdateUser)
}

func main() {
	flag.Parse()

	gin.SetMode(*ginMode)

	r := gin.New()

	// Attach the logger middleware
	r.Use(gin.Logger())

	// Attach the recovery middleware
	r.Use(gin.Recovery())

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
