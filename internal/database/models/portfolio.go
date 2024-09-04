package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Portfolio struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	PortfolioId  string             `bson:"portfolioId"`
	Channel      string             `bson:"channel"`
	Country      string             `bson:"country"`
	CreatedDate  string             `bson:"createdDate"`
	CustomerCode string             `bson:"customerCode"`
	Route        string             `bson:"route"`
}
