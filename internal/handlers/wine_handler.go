package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"wineservice/internal/models"
)

func GetWineById(w http.ResponseWriter, r *http.Request) {
	var wine models.Wine
	wine.ID = "123"
	wine.AlcoholByVolume = 12
	wine.Brand = "henry woj"
	wine.Region = "boston"
	wine.Varietal = "sus"
	wine.Volume = "1"

	err := json.NewEncoder(w).Encode(wine)
	if err != nil {
		fmt.Println(err)
	}

}