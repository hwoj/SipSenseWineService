package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"wineservice/internal/handlers"
	"wineservice/internal/repositories"

	"github.com/gorilla/mux"
)

func StartServer() {
	repositories.ConnectDB()
	defer func() {
		if err := repositories.MongoClient.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	router := mux.NewRouter()

	wineCollection := repositories.MongoClient.Database(os.Getenv("DB_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	wineHandler := handlers.WineHandler{MongoCollection: wineCollection}

	router.HandleFunc("/wine/{id}", wineHandler.GetWineByID).Methods("GET")
	router.HandleFunc("/wine", wineHandler.CreateWine).Methods("POST")
	router.HandleFunc("/wine", wineHandler.GetAllWines).Methods("GET")
	router.HandleFunc("/wine/{id}", wineHandler.DeleteWineByID).Methods("DELETE")
	router.HandleFunc("/wine/{id}", wineHandler.UpdateWineByID).Methods("PUT")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
