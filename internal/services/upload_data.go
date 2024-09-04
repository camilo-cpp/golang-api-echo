package services

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/camilo-cpp/golang-api-echo/internal/database"
	"github.com/camilo-cpp/golang-api-echo/internal/database/models"
)

type UploadDataClient struct {
	db database.MongoConnection
}

func (client *UploadDataClient) UploadData() error {
	portfolioCollection := client.db.Connection().Collection("client-portfolio")
	itemCollection := client.db.Connection().Collection("items-portfolio")

	if checkIfDataExists(portfolioCollection) && checkIfDataExists(itemCollection) {
		return nil
	}

	absPath, err := filepath.Abs("./internal/data/client_portfolio.json")
	if err != nil {
		return err
	}
	portfolios, err := readPortfolios(absPath)
	if err != nil {
		return err
	}

	absPath, err = filepath.Abs("./internal/data/items_portfolio.json")
	if err != nil {
		return err
	}
	items, err := readItems(absPath)
	if err != nil {
		return err
	}

	portfolioResult, err := portfolioCollection.InsertMany(context.TODO(), convertPortfoliosToBson(portfolios))
	if err != nil {
		return err
	}

	portfolioIDMap := make(map[string]primitive.ObjectID)
	for i, portfolio := range portfolios {
		portfolioIDMap[portfolio.ID.Hex()] = portfolioResult.InsertedIDs[i].(primitive.ObjectID)
	}

	modifiedItems := convertItemsToBson(items, portfolioIDMap)

	_, err = itemCollection.InsertMany(context.TODO(), modifiedItems)
	if err != nil {
		return err
	}

	fmt.Println("Data uploaded successfully! 🚀🤘")

	return nil
}

func checkIfDataExists(collection *mongo.Collection) bool {
	count, err := collection.CountDocuments(nil, bson.M{})
	if err != nil {
		fmt.Printf("Error checking data existence: %v\n", err)
		return false
	}
	if count > 0 {
		fmt.Println("Seed data already exists, skipping... 😜")
		return true
	}
	return false
}

func readPortfolios(filename string) ([]models.Portfolio, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var portfolios []models.Portfolio
	if err := json.Unmarshal(bytes, &portfolios); err != nil {
		return nil, err
	}

	return portfolios, nil
}

func readItems(filename string) ([]models.Item, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	var items []models.Item
	if err := json.Unmarshal(bytes, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func convertPortfoliosToBson(portfolios []models.Portfolio) []interface{} {
	var bsonData []interface{}
	for _, portfolio := range portfolios {
		bsonData = append(bsonData, portfolio)
	}
	return bsonData
}

func convertItemsToBson(items []models.Item, portfolioIDMap map[string]primitive.ObjectID) []interface{} {
	var bsonData []interface{}
	for _, item := range items {
		if portfolioID, ok := portfolioIDMap[item.PortfolioID]; ok {
			item.PortfolioID = portfolioID.Hex()
		}
		bsonData = append(bsonData, item)
	}
	return bsonData
}
