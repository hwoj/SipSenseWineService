package repositories

import (
	"context"
	"log"
	"os"
	"wineservice/internal/models"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type WineRepository struct {
	MongoCollection *mongo.Collection
}

var MongoClient *mongo.Client

func ConnectDB() {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	mongoURI := os.Getenv("MONGO_URI")

	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)

	MongoClient, err = mongo.Connect(context.Background(), opts)

	if err != nil {
		log.Fatal("Unable to connect to Mongo", err)
	}

	err = MongoClient.Ping(context.Background(), readpref.Primary())
	if err != nil {
		log.Fatal("ping failed", err)
	}

}

func (wineRepo *WineRepository) GetWineByID(wineID string) (*models.Wine, error) {
	var wine models.Wine

	err := wineRepo.MongoCollection.FindOne(context.Background(), bson.D{{Key: "id", Value: wineID}}).Decode(&wine)

	if err != nil {
		return nil, err
	}

	return &wine, nil
}

func (wineRepo *WineRepository) CreateWine(wine *models.Wine) (interface{}, error) {
	result, err := wineRepo.MongoCollection.InsertOne(context.Background(), wine)

	if wine != nil {
		return nil, err
	}

	return result.InsertedID, nil
}

func (wineRepo *WineRepository) GetAllWines() ([]models.Wine, error) {
	results, err := wineRepo.MongoCollection.Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var wines []models.Wine
	err = results.All(context.Background(), &wines)

	if err != nil {
		return nil, err
	}
	return wines, nil
}

func (wineRepo *WineRepository) UpdateWine(wineID string, updateWine *models.Wine) (int64, error) {
	result, err := wineRepo.MongoCollection.UpdateOne(context.Background(), bson.D{{Key: "id", Value: wineID}}, bson.D{{Key: "$set", Value: updateWine}})

	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (wineRepo *WineRepository) DeleteByID(wineID string) (int64, error) {
	result, err := wineRepo.MongoCollection.DeleteOne(context.Background(), bson.D{{Key: "id", Value: wineID}})

	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
