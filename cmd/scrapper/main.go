package main

func main() {

	stations := new(Stations)
	scrapVelobikePage(stations)
	traceStations(stations)

	println("Scrapper is starting")
}
