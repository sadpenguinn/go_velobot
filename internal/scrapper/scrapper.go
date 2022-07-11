package scrapper

import (
	"encoding/json"
	"log"
	"net/http"
)

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
