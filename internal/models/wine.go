package models

type Wine struct {
	ID              string `json:"id"`
	Brand           string `json:"brand"`
	Varietal        string `json:"varietal"`
	Region          string `json:"region"`
	Volume          string `json:"volume"`
	AlcoholByVolume int    `json:"alcoholByVolume"`
}
