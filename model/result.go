package model

import (
	"encoding/json"
)

type WeatherResponse struct {
	FeelsLike float64 // takes main.feels_like
	TempMin   float64 // takes main.temp_min
	TempMax   float64 // takes main.temp_max
}

func (rsp *WeatherResponse) UnmarshalJSON(b []byte) error {

	var f interface{}
	if err := json.Unmarshal(b, &f); err != nil {
		return err
	}

	m := f.(map[string]interface{})

	temp := m["main"]
	v := temp.(map[string]interface{})

	rsp.FeelsLike = v["feels_like"].(float64)
	rsp.TempMin = v["temp_min"].(float64)
	rsp.TempMax = v["temp_max"].(float64)

	return nil
}
