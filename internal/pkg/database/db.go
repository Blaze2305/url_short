package database

import "github.com/Blaze2305/url_short/internal/pkg/model"

const databaseName = "shortner"

// DbMethods - Interface to help maintain and implement all DB methods
type DbMethods interface {
	ListUrls() (*[]model.Shorten, error)
	CreateURL(token model.Shorten) (*model.Shorten, error)
	GetURL(token string) (*string, error)
	DeleteURL(token string) (*string, error)

	CreateUser(input model.User) (*model.User, error)
	GetUser(uid string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	DeleteUser(uid string) (*string, error)
	UpdateUser(input model.User) error

	CreateToken(input model.Token) (*model.Token, error)
	DeleteToken(token string) (*string, error)
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
