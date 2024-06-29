package api

import (
	"log"
	"net/http"
	"wineservice/internal/handlers"

	"github.com/gorilla/mux"
)

func StartServer() {
	router := mux.NewRouter()
	router.HandleFunc("/wine/{id}", handlers.GetWineById).Methods("GET")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
