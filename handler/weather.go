package handler

import (
	"encoding/json"
	"fmt"
	model "github.com/alokgarg-dev/weather-api/model"
	client "github.com/alokgarg-dev/weather-api/service"
	utils "github.com/alokgarg-dev/weather-api/utils"
	"log"
	"net/http"
	"strconv"
)

type Weather struct{}

func (wt *Weather) prepareOpenWeatherUrl(param1 string, param2 string) (string, error) {

	const bitSize = 64

	latitude, err := strconv.ParseFloat(param1, bitSize)
	if err != nil {
		return "", err
	}

	longitude, err := strconv.ParseFloat(param2, bitSize)
	if err != nil {
		return "", err
	}

	openWeatherUrl := utils.CreateUrl(latitude, longitude)
	return openWeatherUrl, nil
}

func (wt *Weather) prepareWeatherResponse(response *model.WeatherResponse, param1 string, param2 string) string {
	var currentWeather string = "Weather at latitude " + param1 + " and longitude " + param2 + " is currently"
	if response.FeelsLike < 15 {
		currentWeather += " Cold!!"
	} else if response.FeelsLike >= 15.0 && response.FeelsLike < 30.0 {
		currentWeather += " Moderate"
	} else {
		currentWeather += " Hot!!"
	}
	return currentWeather
}

func (wt *Weather) Get(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get Weather Details")
	// Parse query parameters
	params := r.URL.Query()
	param1 := params.Get("lat")
	param2 := params.Get("lon")

	openWeatherUrl, err := wt.prepareOpenWeatherUrl(param1, param2)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Request received"))
		return
	}

	body, err := client.SendRequest(openWeatherUrl)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Json Response from OpenWeather Api"))
		return
	}

	var response *model.WeatherResponse
	// Unmarshal or Decode the JSON to the interface.
	if err = json.Unmarshal([]byte(body), &response); err != nil {
		log.Println(err)
		w.Write([]byte("Unable to parse Json Response from OpenWeather Api"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	currentWeather := wt.prepareWeatherResponse(response, param1, param2)
	w.Write([]byte(currentWeather))
	w.WriteHeader(http.StatusOK)
}
