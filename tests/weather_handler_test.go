package tests

import (
	"cloud-run-weather-api/handlers"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInvalidCEP(t *testing.T) {
	req := httptest.NewRequest("GET", "/weather?cep=abc123", nil)
	w := httptest.NewRecorder()

	handlers.GetWeatherHandler(w, req)

	if w.Code != http.StatusUnprocessableEntity {
		t.Errorf("Expected 422, got %d", w.Code)
	}
}
