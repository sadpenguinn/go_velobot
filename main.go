package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var redisClient = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "",
	DB:       0,
})

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
	//rememberStations(stations)
	//traceStationById("0969")
}
