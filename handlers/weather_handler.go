package handlers

import (
	"encoding/json"
	"log"
	"math"
	"net/http"
	"regexp"

	"cloud-run-weather-api/services"
	"cloud-run-weather-api/utils"
)

func round(value float64, precision int) float64 {
	mult := math.Pow(10, float64(precision))
	return math.Round(value*mult) / mult
}

func GetWeatherHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Received /weather request")
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		log.Println("Error: cep parameter is missing")
		http.Error(w, "Missing cep parameter", http.StatusBadRequest)
		return
	}
	log.Println("CEP recebido:", cep)

	if !regexp.MustCompile(`^\d{8}$`).MatchString(cep) {
		log.Println("Error: invalid zipcode format")
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	location, err := services.GetLocationByCEP(cep)
	if err != nil {
		log.Printf("Error fetching location by CEP: %v\n", err)
		http.Error(w, "error looking up zipcode", http.StatusInternalServerError)
		return
	}
	if location.Localidade == "" {
		log.Println("Error: location not found for CEP")
		http.Error(w, "can not find zipcode", http.StatusNotFound)
		return
	}
	log.Printf("Location found: %v\n", location.Localidade)

	tempC, err := services.GetTemperature(location.Localidade)
	if err != nil {
		log.Printf("Error fetching temperature: %v\n", err)
		http.Error(w, "weather service error", http.StatusInternalServerError)
		return
	}
	log.Printf("Temperature fetched: %.2f C\n", tempC)

	response := map[string]float64{
		"temp_C": round(tempC, 2),
		"temp_F": round(utils.CelsiusToFahrenheit(tempC), 2),
		"temp_K": round(utils.CelsiusToKelvin(tempC), 2),
	}

	respJson, err := json.Marshal(response)
	if err != nil {
		log.Printf("Error encoding JSON response: %v\n", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	log.Println("Sending successful response:", string(respJson))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(respJson)
}
