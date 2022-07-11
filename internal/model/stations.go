package model

type Stations struct {
	Items []struct {
		Address                string `json:"Address"`
		AvailableElectricBikes int    `json:"AvailableElectricBikes"`
		AvailableOrdinaryBikes int    `json:"AvailableOrdinaryBikes"`
		FreeElectricPlaces     int    `json:"FreeElectricPlaces"`
		FreeOrdinaryPlaces     int    `json:"FreeOrdinaryPlaces"`
		FreePlaces             int    `json:"FreePlaces"`
		HasTerminal            bool   `json:"HasTerminal"`
		IconsSet               int    `json:"IconsSet"`
		ID                     string `json:"Id"`
		IsFavourite            bool   `json:"IsFavourite"`
		IsLocked               bool   `json:"IsLocked"`
		IsOverflow             bool   `json:"IsOverflow"`
		Name                   string `json:"Name"`
		Position               struct {
			Lat float64 `json:"Lat"`
			Lon float64 `json:"Lon"`
		} `json:"Position"`
		StationTypes        []string `json:"StationTypes"`
		TemplateID          int      `json:"TemplateId"`
		TotalElectricPlaces int      `json:"TotalElectricPlaces"`
		TotalOrdinaryPlaces int      `json:"TotalOrdinaryPlaces"`
		TotalPlaces         int      `json:"TotalPlaces"`
	} `json:"Items"`
}
