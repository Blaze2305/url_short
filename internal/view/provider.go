package view

import (
	"github.com/Blaze2305/url_short/internal/pkg/database"
	"github.com/gin-gonic/gin"
)

// ProviderMethods - Provides an interface to keep code clean
type ProviderMethods interface {
	CreateURL(c *gin.Context)
	// ListURLs(c *gin.Context)
	Redirect(c *gin.Context)
}

// Provider - connection between db and view
type Provider struct {
	db database.DbMethods
}

// NewProvider - Create a  New provider to use
func NewProvider(db database.DbMethods) ProviderMethods {
	return Provider{db: db}
}
