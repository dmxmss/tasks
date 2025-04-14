package repository

import (
	"github.com/dmxmss/tasks/config"
	"github.com/dmxmss/tasks/entities"
	e "github.com/dmxmss/tasks/error"

	"fmt"
	"net/http"
	"encoding/json"
	"strings"
)

type WeatherRepository interface {
	GetCurrentWeatherAt(city string) (*entities.WeatherResponse, error) 
}

type weatherRepository struct {
	conf *config.Weather
}

func NewWeatherRepository(conf *config.Weather) WeatherRepository {
	url := "https://api.weatherapi.com/v1/current.json?q=%s&key=%s"
	conf.URL = url
	return &weatherRepository{
		conf: conf,
	}
}

func (w *weatherRepository) GetCurrentWeatherAt(city string) (*entities.WeatherResponse, error) {
	req := fmt.Sprintf(w.conf.URL, strings.ReplaceAll(city, " ", "%20"), w.conf.Key)

	response, err := http.Get(req)	
	if err != nil {
		return nil, e.ErrGetWeatherFailed
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return nil, e.ErrCityNotFound
	}

	var weatherResponse entities.WeatherResponse
	if err := json.NewDecoder(response.Body).Decode(&weatherResponse); err != nil {
		return nil, e.ErrGetWeatherFailed
	}

	return &weatherResponse, nil
}
