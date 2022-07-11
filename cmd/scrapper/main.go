package main

import "github.com/sadpenguinn/go_velobot/internal/model"

func main() {

	stations := new(Stations)
	scrapVelobikePage(stations)
	traceStations(stations)

	println("Scrapper is starting")
}
