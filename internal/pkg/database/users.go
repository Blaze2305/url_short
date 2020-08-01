package database

import (
	"context"
	"time"

	"github.com/Blaze2305/url_short/internal/pkg/constants"
	"github.com/Blaze2305/url_short/internal/pkg/model"
	logger "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateUser - create a user
func (d db) CreateUser(input model.User) (*model.User, error) {
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

	coll := client.Database(d.dbName).Collection(constants.UserCollection)

	_, err = coll.InsertOne(ctx, input)
	if err != nil {
		logger.Errorf("Error during insertion %s", err.Error())
		return nil, err
	}
	return &input, nil
}

// GetUser - get user details
func (d db) GetUser(uid string) (*model.User, error) {
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

	coll := client.Database(d.dbName).Collection(constants.UserCollection)

	user := model.User{}

	cur := coll.FindOne(ctx, bson.M{"_id": uid})

	if err := cur.Decode(&user); err != nil {
		logger.Errorf("db: error while getting url %s", err.Error())
		return nil, err
	}
	return &user, nil
}

// DeleteUser - delete a user
func (d db) DeleteUser(uid string) (*string, error) {
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

	coll := client.Database(d.dbName).Collection(constants.UserCollection)

	resp, err := coll.DeleteOne(ctx, bson.M{"_id": uid})
	if err != nil {
		logger.Errorf("error while deleting user %s", err.Error())
		return nil, err
	}

	if resp.DeletedCount != 1 {
		logger.Errorf("did not delete user %s", uid)
		return nil, err
	}

	urlColl := client.Database(d.dbName).Collection(constants.URLCollection)

	_, err = urlColl.DeleteMany(ctx, bson.M{"user": uid})
	if err != nil {
		logger.Errorf("error while deleting user %s", err.Error())
		return nil, err
	}
	return &uid, nil
}
