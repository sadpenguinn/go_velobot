package log

import (
	"encoding/json"
)

func traceStations(stations *Stations) {
	for _, station := range stations.Items {
		stationJson, _ := json.Marshal(station)
		println(string(stationJson) + "\n")
	}
}
