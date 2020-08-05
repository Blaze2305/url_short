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
	count, err := coll.CountDocuments(ctx, bson.M{"userid": input.UserID})
	if err != nil {
		logger.Errorf("Error during insertion %s", err.Error())
		return nil, err
	}

	if count > 0 {
		_, err = coll.DeleteOne(ctx, bson.M{"userid": input.UserID})
		if err != nil {
			logger.Errorf("Error during insertion %s", err.Error())
			return nil, err
		}
	}
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
	if result.DeletedCount == 0 {
		return nil, errors.New("User already logged out")
	}

	if result.DeletedCount < 1 {
		logger.Errorf("Unable to delete token from database")
		return nil, errors.New("Unable to delete token from database")
	}

	return &token, nil
}

// GetUserFromToken - get the user model using the token
func (d db) GetUserFromToken(token string) (*model.User, error) {
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

	userColl := client.Database(d.dbName).Collection(constants.UserCollection)

	tokenObj := model.Token{}
	resp := coll.FindOne(ctx, bson.M{"_id": token})
	if err := resp.Decode(&tokenObj); err != nil {
		logger.Errorf("error while fetching token %s", err.Error())
		return nil, err
	}

	UserObj := model.User{}
	resp = userColl.FindOne(ctx, bson.M{"_id": tokenObj.UserID})
	if err = resp.Decode(&UserObj); err != nil {
		logger.Errorf("error while fetching user %s", err.Error())
		return nil, err
	}

	return &UserObj, nil
}
