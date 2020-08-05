package database

import (
	"context"
	"errors"
	"time"

	"github.com/Blaze2305/url_short/internal/pkg/constants"
	"github.com/Blaze2305/url_short/internal/pkg/model"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateToken - creates token when user login
func (d db) CreateToken(input model.Token) (*model.Token, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(d.connection))
	if err != nil {
		logger.Errorf("error while connecting to db %s", err.Error())
		return nil, err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			logger.Errorf("error while disconnecting db %s", err.Error())
		}
	}()

	coll := client.Database(d.dbName).Collection(constants.TokenCollection)

	_, err = coll.InsertOne(ctx, input)
	if err != nil {
		logger.Errorf("Error during insertion %s", err.Error())
		return nil, err
	}
	return &input, nil
}

// DeleteToken - creates token when user login
func (d db) DeleteToken(token string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(d.connection))
	if err != nil {
		logger.Errorf("error while connecting to db %s", err.Error())
		return nil, err
	}

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			logger.Errorf("error while disconnecting db %s", err.Error())
		}
	}()

	coll := client.Database(d.dbName).Collection(constants.TokenCollection)

	result, err := coll.DeleteOne(ctx, bson.M{"_id": token})
	if err != nil {
		logger.Errorf("error while connecting to db %s", err.Error())
		return nil, err
	}

	if result.DeletedCount < 1 {
		logger.Errorf("Unable to delete token from database")
		return nil, errors.New("Unable to delete token from database")
	}

	return &token, nil
}
