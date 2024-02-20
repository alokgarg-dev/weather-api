package utils

import (
	"github.com/spf13/viper"
	"strconv"
)

func CreateUrl(latitude, longitude float64) string {
	openWeatherUrl := "https://api.openweathermap.org/data/2.5/weather?lat="
	openWeatherUrl += strconv.FormatFloat(latitude, 'f', -1, 64)
	openWeatherUrl += "&lon=" + strconv.FormatFloat(longitude, 'f', -1, 64)
	openWeatherUrl += "&appid=" + viper.GetString("API_KEY")
	openWeatherUrl += "&units=metric"
	return openWeatherUrl
}
