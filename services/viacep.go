package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"cloud-run-weather-api/models"
)

func GetLocationByCEP(cep string) (*models.ViaCEPResponse, error) {
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	log.Println("[DEBUG] Corpo da resposta do ViaCEP:", string(bodyBytes))

	var location models.ViaCEPResponse
	if err := json.Unmarshal(bodyBytes, &location); err != nil {
		log.Println("[ERROR] Falha ao decodificar JSON:", err)
		return nil, err
	}

	if location.Localidade == "" {
		log.Println("[ERROR] Localidade vazia para o CEP:", cep)
		return nil, errors.New("can not find zipcode")
	}

	log.Println("[DEBUG] Localidade encontrada:", location.Localidade)
	return &location, nil
}
