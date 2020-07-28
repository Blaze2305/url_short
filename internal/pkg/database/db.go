package database

import "github.com/Blaze2305/url_short/internal/pkg/model"

const databaseName = "shortner"

// DbMethods - Interface to help maintain and implement all DB methods
type DbMethods interface {
	ListUrls() (*[]model.Shorten, error)
	CreateURL(token model.Shorten) (*model.Shorten, error)
}

type db struct {
	connection string
	dbName     string
}

// NewDatabase - a New Database object to handle db methods
func NewDatabase(connection string) DbMethods {
	return db{
		connection: connection,
		dbName:     databaseName,
	}
}
