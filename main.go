package main

import (
	"log"
	"net/http"
	"os"

	"cloud-run-weather-api/handlers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: .env file not found or failed to load")
	}

	log.Println("WEATHER_API_KEY =", os.Getenv("WEATHER_API_KEY"))

	http.HandleFunc("/weather", handlers.GetWeatherHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Running on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
