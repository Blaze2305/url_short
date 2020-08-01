package view

import (
	"github.com/Blaze2305/url_short/internal/pkg/database"
	"github.com/gin-gonic/gin"
)

// ProviderMethods - Provides an interface to keep code clean
type ProviderMethods interface {
	Redirect(c *gin.Context)

	CreateURL(c *gin.Context)
	ListURLs(c *gin.Context)
	DeleteURL(c *gin.Context)

	CreateUser(c *gin.Context)
	GetUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUser(c *gin.Context)
}

// Provider - connection between db and view
type Provider struct {
	db database.DbMethods
}

// NewProvider - Create a  New provider to use
func NewProvider(db database.DbMethods) ProviderMethods {
	return Provider{db: db}
}
