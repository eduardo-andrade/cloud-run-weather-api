package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

type WeatherAPIResponse struct {
	Current struct {
		TempC float64 `json:"temp_c"`
	} `json:"current"`
}

func GetTemperature(city string) (float64, error) {
	apiKey := os.Getenv("WEATHER_API_KEY")

	// Escapa o nome da cidade para URL válida
	cityEncoded := url.QueryEscape(city)

	url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=%s&q=%s", apiKey, cityEncoded)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("weather API returned status: %s", resp.Status)
	}

	var weather WeatherAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&weather); err != nil {
		return 0, err
	}

	return weather.Current.TempC, nil
}
