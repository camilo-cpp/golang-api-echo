package services

import (
	"context"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/camilo-cpp/golang-api-echo/internal/database"
	"github.com/camilo-cpp/golang-api-echo/internal/database/models"
	"github.com/camilo-cpp/golang-api-echo/internal/dtos"
)

type GetPortfolioByClientIdClient struct {
	db database.MongoConnection
}

func (client *GetPortfolioByClientIdClient) GetPortfolioByClientIdService(clientId string) (*models.Portfolio, error) {
	collection := client.db.Connection().Collection("client-portfolio")

	filter := bson.M{
		"customerCode": clientId,
	}

	var portfolio models.Portfolio

	err := collection.FindOne(context.Background(), filter).Decode(&portfolio)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return &portfolio, nil
}

func (client *GetPortfolioByClientIdClient) GetPortfolioItemsByClientIdService(params *dtos.ParamsGetPortfolioItemsByClientId) ([]*models.Item, error) {
	collection := client.db.Connection().Collection("items-portfolio")

	page, err := strconv.Atoi(params.CurrentPage)
	if err != nil {
		return nil, err
	}

	limit, err := strconv.Atoi(params.PageSize)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"portfolioId": params.PortfolioId,
	}

	aggregation := []bson.D{}
	aggregation = append(aggregation, bson.D{{Key: "$match", Value: filter}})
	aggregation = append(aggregation, bson.D{{Key: "$skip", Value: (page - 1)}})
	aggregation = append(aggregation, bson.D{{Key: "$limit", Value: limit}})

	var items []*models.Item

	cursor, err := collection.Aggregate(context.Background(), aggregation)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var item models.Item
		if err := cursor.Decode(&item); err != nil {
			return nil, err
		}
		items = append(items, &item)
	}

	return items, nil
}
