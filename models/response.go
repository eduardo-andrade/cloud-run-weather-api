package models

import "encoding/json"

type ViaCEPResponse struct {
	Cep         string          `json:"cep"`
	Logradouro  string          `json:"logradouro"`
	Complemento string          `json:"complemento"`
	Unidade     string          `json:"unidade"`
	Bairro      string          `json:"bairro"`
	Localidade  string          `json:"localidade"`
	Uf          string          `json:"uf"`
	Estado      string          `json:"estado"`
	Regiao      string          `json:"regiao"`
	Ibge        string          `json:"ibge"`
	Gia         string          `json:"gia"`
	Ddd         string          `json:"ddd"`
	Siafi       string          `json:"siafi"`
	RawErro     json.RawMessage `json:"erro,omitempty"`
}

func (v *ViaCEPResponse) HasError() bool {
	if len(v.RawErro) == 0 {
		return false
	}
	var b bool
	if err := json.Unmarshal(v.RawErro, &b); err == nil {
		return b
	}
	var s string
	if err := json.Unmarshal(v.RawErro, &s); err == nil {
		return s == "true"
	}
	return false
}
