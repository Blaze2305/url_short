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

// CreateURL - Create a short url for the given input url
func (d db) CreateURL(input model.Shorten) (*model.Shorten, error) {
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

	coll := client.Database(d.dbName).Collection(constants.URLCollection)

	_, err = coll.InsertOne(ctx, input)
	if err != nil {
		logger.Errorf("Error during insertion %s", err.Error())
		return nil, err
	}
	return &input, nil
}

// ListUrls - List all available shortned urls
func (d db) ListUrls() (*[]model.Shorten, error) {
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

	coll := client.Database(d.dbName).Collection(constants.URLCollection)

	Urls := []model.Shorten{}

	curr, err := coll.Find(ctx, bson.M{})

	if err = curr.All(ctx, &Urls); err != nil {
		logger.Errorf("db: error while list urls %s", err.Error())
		return nil, err
	}

	return &Urls, nil
}

func (d db) GetURL(token string) (*string, error) {
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

	coll := client.Database(d.dbName).Collection(constants.URLCollection)

	out := model.Shorten{}

	curr := coll.FindOne(ctx, bson.M{"_id": token})

	if err = curr.Decode(&out); err != nil {
		logger.Errorf("db: error while getting url %s", err.Error())
		return nil, err
	}

	return &out.Forward, nil
}
