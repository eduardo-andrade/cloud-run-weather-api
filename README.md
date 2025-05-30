
# Projeto Weather API

Este projeto em Go é uma API que retorna a temperatura atual de uma cidade brasileira com base em um CEP informado. A aplicação consulta o [ViaCEP](https://viacep.com.br/) para identificar a cidade correspondente ao CEP e depois consome a [WeatherAPI](https://www.weatherapi.com/) para obter a temperatura atual.

A aplicação está disponível em produção via **Cloud Run**:

🔗 [https://cloudrun-weather-api-104778723545.us-central1.run.app/weather?cep=01001000](https://cloudrun-weather-api-104778723545.us-central1.run.app/weather?cep=01001000)

---

## Estrutura de Diretórios

```
.
├── Dockerfile
├── go.mod
├── go.sum
├── handlers
│   └── weather_handler.go
├── LICENSE
├── main.go
├── models
│   └── response.go
├── services
│   ├── viacep.go
│   └── weatherapi.go
├── tests
│   └── weather_handler_test.go
└── utils
    └── convert.go
```

---

## Requisitos

- Go 1.21 ou superior
- Docker
- Chave da [WeatherAPI](https://www.weatherapi.com/)

---

## Instalação

Clone o repositório e acesse a pasta do projeto:

```bash
git clone <url-do-repositorio>
cd cloud-run-weather-api
```

Baixe as dependências:

```bash
go mod tidy
```

---

## Configuração

Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:

```
WEATHER_API_KEY=your_weatherapi_key
```

Substitua `your_weatherapi_key` pela sua chave real da WeatherAPI.

---

## Execução

### Executando Localmente

```bash
go run main.go
```

Depois, acesse:  
[http://localhost:8080/weather?cep=01001000](http://localhost:8080/weather?cep=01001000)

---

### Executando com Docker

```bash
docker build -t cloudrun-weather-api .
docker run --env-file .env -p 8080:8080 cloudrun-weather-api
```

---

## Testes

### Via Navegador ou Postman

Acesse diretamente no navegador ou use ferramentas como Postman com a seguinte URL:

```
http://localhost:8080/weather?cep=01001000
```

### Via `curl` no Terminal

Você também pode testar a API com o `curl` no terminal:

```bash
curl "http://localhost:8080/weather?cep=01001000"
```

Ou, caso deseje testar a versão em produção no Cloud Run:

```bash
curl "https://cloudrun-weather-api-104778723545.us-central1.run.app/weather?cep=01001000"
```

> O parâmetro `cep` é obrigatório e pode ser substituído por qualquer CEP válido do Brasil.  
> Exemplo:
> ```bash
> curl "https://cloudrun-weather-api-104778723545.us-central1.run.app/weather?cep=30140071"
> ```

---

## Resposta Esperada

```json
{
  "temp_C": 14.1,
  "temp_F": 57.38,
  "temp_K": 287.1
}
```

Os valores representam:

- `temp_C`: Temperatura em Celsius  
- `temp_F`: Temperatura em Fahrenheit  
- `temp_K`: Temperatura em Kelvin  

---

## Possíveis Erros e Soluções

### ❌ Erro: `Missing cep parameter`

- O parâmetro `cep` não foi informado na URL.  
  Exemplo correto: `/weather?cep=01001000`

### ❌ Erro: `invalid zipcode`

- O CEP deve conter **8 dígitos numéricos**, sem traços ou espaços.

### ❌ Erro: `can not find zipcode`

- O CEP informado não existe ou não foi encontrado pela API do ViaCEP.

### ❌ Erro: `weather service error`

- Pode indicar:
  - Falha na consulta à WeatherAPI
  - `WEATHER_API_KEY` ausente ou inválida
  - Excesso de requisições (limite da API)

---

## Considerações Finais

Este projeto demonstra a construção de uma API REST com Go integrando múltiplas fontes de dados externas. É útil para aprendizado sobre manipulação de parâmetros, requisições HTTP, estrutura modular em Go e deployment com Docker na Google CloudRun.