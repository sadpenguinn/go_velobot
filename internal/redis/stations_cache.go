package redis

import (
	"encoding/json"
	"log"
)

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
