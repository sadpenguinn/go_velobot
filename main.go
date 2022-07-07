package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"net/http"
)

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

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

func scrapVelobikePage(data interface{}) {
	// Build request
	request, err := http.NewRequest("OPTIONS", "https://velobike.ru/ajax/parkings", nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Cookie", "qrator_jsr=1657202838.110.PiJfOZ6OPUwQjcOf-nkojg6c2u0u9hg1luqllsiipm1t7i58r-00; qrator_jsid=1657202838.110.PiJfOZ6OPUwQjcOf-1umh2h1flvb6b2ekssa3l5fmvobmt9qt")

	// Send request, get response
	httpClient := &http.Client{}
	response, err := httpClient.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Validate response
	if response.StatusCode != 200 {
		log.Fatal("Invalid status: " + response.Status)
	}

	// Parse response
	err = json.NewDecoder(response.Body).Decode(data)
	if err != nil {
		log.Fatal(err)
	}
}

func rememberStations(stations *Stations) {
	for _, station := range stations.Items {
		stationJson, err := json.Marshal(station)
		if err != nil {
			log.Fatal(err)
		}
		err = redisClient.Set(station.ID, stationJson, 0).Err()
		if err != nil {
			log.Fatal(err)
		}
		println(station.ID)
	}
}

func traceStations(stations *Stations) {
	for _, station := range stations.Items {
		stationJson, _ := json.Marshal(station)
		println(string(stationJson) + "\n")
	}
}

func traceStationById(id string) {
	val, err := redisClient.Get(id).Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val)
}

func main() {
	stations := new(Stations)
	scrapVelobikePage(stations)
	traceStations(stations)
	rememberStations(stations)
	traceStationById("0969")
}
