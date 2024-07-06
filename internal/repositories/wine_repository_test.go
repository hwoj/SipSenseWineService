package repositories

import (
	"context"
	"log"
	"os"
	"testing"
	"wineservice/internal/models"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	MongoURI := os.Getenv("MONGO_URI")

	mongoTestClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI(MongoURI))

	if err != nil {
		log.Fatalln("Failed to connect to mongodb server", err)
	}

	log.Println("Successfully connected to mongodb atlas")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatalln("Ping to mongo failed", err)
	}

	log.Println("Ping successful")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	// dummy wine data
	wineID1 := uuid.New().String()
	wineID2 := uuid.New().String()

	var wine models.Wine

	// connect to collection

	wines_test_collection := mongoTestClient.Database("wine").Collection("wines_test")

	wineRepository := WineRepository{MongoCollection: wines_test_collection}

	// insert employee
	t.Run("Insert Wine 1", func(t *testing.T) {
		wine = models.Wine{
			ID:              wineID1,
			Brand:           "Henry Cellars",
			Varietal:        "Cabernet Sauvignon",
			Region:          "Boston, MA",
			Volume:          "750ml",
			AlcoholByVolume: 15,
			Image:           "bob.com",
		}

		result, err := wineRepository.CreateWine(&wine)

		if err != nil {
			t.Fatal("Failed to insert wine 1: ", err)
		}

		t.Log("Insert Wine 1 successful", result)

	})

	t.Run("Insert Wine 2", func(t *testing.T) {
		wine = models.Wine{
			ID:              wineID2,
			Brand:           "Ella Cellars",
			Varietal:        "Chardonnay",
			Region:          "California",
			Volume:          "750ml",
			AlcoholByVolume: 17.5,
			Image:           "bob.com",
		}

		result, err := wineRepository.CreateWine(&wine)

		if err != nil {
			t.Fatal("Failed to insert wine 2:", err)
		}

		t.Log("Insert Wine 2 successful:", result)

	})

	t.Run("Get Wine 1", func(t *testing.T) {
		wine, err := wineRepository.GetWineByID(wineID1)

		if err != nil {
			t.Fatal("Failed to get wine 1:", err)
		}

		t.Log("Get Wine 1 successful:", wine.Brand)
	})

	t.Run("Get All Wines", func(t *testing.T) {
		wines, err := wineRepository.GetAllWines()

		if err != nil {
			t.Fatal("Failed to get all wines:", err)
		}

		t.Log("Get All Wines successful:", wines)
	})

	t.Run("Update Wine 1 Brand", func(t *testing.T) {
		wine = models.Wine{
			ID:              wineID1,
			Brand:           "Wojnicki Cellars",
			Varietal:        "Cabernet Sauvignon",
			Region:          "Boston, MA",
			Volume:          "750ml",
			AlcoholByVolume: 15,
			Image:           "bob.com",
		}

		updateCount, err := wineRepository.UpdateWine(wineID1, &wine)

		if err != nil {
			t.Fatal("Failed to update wine 1 brand:", err)
		}

		t.Log("Update Wine 1 Brand successful. Update count:", updateCount)
	})

	t.Run("Get Wine 1 After Update", func(t *testing.T) {
		wine, err := wineRepository.GetWineByID(wineID1)

		if err != nil {
			t.Fatal("Failed to get wine 1 after update:", err)
		}

		t.Log("Get Wine 1 After Update successful:", wine.Brand)
	})

	t.Run("Delete Wine 2", func(t *testing.T) {
		deleteCount, err := wineRepository.DeleteByID(wineID2)

		if err != nil {
			t.Fatal("Delete Wine 2 failed:", err)
		}

		t.Log("Delete Wine 2 successful. Delete count:", deleteCount)
	})

	t.Run("Get All Wines After Deletion", func(t *testing.T) {
		wines, err := wineRepository.GetAllWines()

		if err != nil {
			t.Fatal("Get All Wines After Deletion failed:", err)
		}

		t.Log("Get All Wines After Deletion successful: ", wines)
	})
}
