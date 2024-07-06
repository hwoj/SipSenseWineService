package models

type Wine struct {
	ID              string  `json:"id,omitempty" bson:"id"`
	Brand           string  `json:"brand,omitempty" bson:"brand"`
	Varietal        string  `json:"varietal,omitempty" bson:"varietal"`
	Region          string  `json:"region,omitempty" bson:"region"`
	Volume          string  `json:"volume,omitempty" bson:"volume"`
	AlcoholByVolume float32 `json:"alcoholByVolume,omitempty" bson:"alcoholByVolume"`
	Image           string  `json:"image,omitempty" bson:"image"`
	SugarContent    string  `json:"sugarContent,omitempty" bson:"sugarContent"`
}
