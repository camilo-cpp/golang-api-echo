package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Item struct {
	ID                     primitive.ObjectID `bson:"_id,omitempty"`
	PortfolioID            string             `bson:"portfolioId"`
	SKU                    string             `bson:"sku"`
	Title                  string             `bson:"title"`
	CategoryID             string             `bson:"categoryId"`
	Category               string             `bson:"category"`
	Brand                  string             `bson:"brand"`
	Classification         string             `bson:"classification"`
	UnitsPerBox            string             `bson:"unitsPerBox"`
	MinOrderUnits          string             `bson:"minOrderUnits"`
	PackageDescription     string             `bson:"packageDescription"`
	PackageUnitDescription string             `bson:"packageUnitDescription"`
	QuantityMaxRedeem      int                `bson:"quantityMaxRedeem"`
	RedeemUnit             string             `bson:"redeemUnit"`
	OrderReasonRedeem      string             `bson:"orderReasonRedeem"`
	SKURedeem              bool               `bson:"skuRedeem"`
	Price                  Price              `bson:"price"`
	Points                 int                `bson:"points"`
}
