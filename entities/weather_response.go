package entities

import (
	"encoding/json"
)

type WeatherResponse struct {
	Current struct { 
		Temperature float32 `json:"temp_c"`
		Condition struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
		} `json:"condition"`
	} `json:"current"`
}

func (w WeatherResponse) String() string {
	bytes, _ := json.Marshal(map[string]any{
		"temp": w.Current.Temperature,
		"desc": w.Current.Condition.Text,
		"icon": w.Current.Condition.Icon,
	})

	return string(bytes)
}
