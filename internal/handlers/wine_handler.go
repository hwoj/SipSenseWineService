package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"wineservice/internal/models"
	"wineservice/internal/repositories"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type WineHandler struct {
	MongoCollection *mongo.Collection
}

func (handler WineHandler) GetWineByID(w http.ResponseWriter, r *http.Request) {

	wineID := mux.Vars(r)["id"]

	wineRepository := repositories.WineRepository{MongoCollection: handler.MongoCollection}

	wine, err := wineRepository.GetWineByID(wineID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("unable to retrieve wine:", err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(wine)
	if err != nil {
		fmt.Println(err)
	}

}

func (handler WineHandler) CreateWine(w http.ResponseWriter, r *http.Request) {
	var wine models.Wine
	err := json.NewDecoder(r.Body).Decode(&wine)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("invalid request body:", err)
	}

	wine.ID = uuid.NewString()

	wineRepository := repositories.WineRepository{MongoCollection: handler.MongoCollection}

	insertID, err := wineRepository.CreateWine(&wine)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("unable to create new wine:", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, insertID)

	log.Println("Wine created with ID", insertID)

}

func (handler WineHandler) GetAllWines(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	wineRepository := repositories.WineRepository{MongoCollection: handler.MongoCollection}
	wines, err := wineRepository.GetAllWines()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Unable to retrieve all wines:", err)
	}

	err = json.NewEncoder(w).Encode(wines)
	if err != nil {
		fmt.Println(err)
	}
}

func (handler WineHandler) DeleteWineByID(w http.ResponseWriter, r *http.Request) {
	wineID := mux.Vars(r)["id"]
	wineRepository := repositories.WineRepository{MongoCollection: handler.MongoCollection}

	deletedCount, err := wineRepository.DeleteByID(wineID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("unable to delete wine:", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted %d wines", deletedCount)

}

func (handler WineHandler) UpdateWineByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	wineID := mux.Vars(r)["id"]

	var wine models.Wine
	err := json.NewDecoder(r.Body).Decode(&wine)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Invalid request body:", err)
	}

	wine.ID = wineID

	wineRepository := repositories.WineRepository{MongoCollection: handler.MongoCollection}
	updateCount, err := wineRepository.UpdateWine(wineID, &wine)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("unable to update wine:", err)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "updated %d wine", updateCount)

}
