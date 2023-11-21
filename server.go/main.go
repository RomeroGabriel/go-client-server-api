package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	integration "github.com/RomeroGabriel/server.go/integration"
	service "github.com/RomeroGabriel/server.go/service"
)

func ConvertHighHandler(res http.ResponseWriter, req *http.Request) {
	api := integration.NewExchangeRateApi("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	service := service.NewQuotationService(*api)

	queryData := req.URL.Query()
	if !queryData.Has("value") {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	value_par := req.URL.Query().Get("value")
	value, err := strconv.ParseFloat(value_par, 64)
	if err != nil {
		// handle error
		res.WriteHeader(http.StatusBadRequest)
		return
	}

	result, err := service.CalculateRealToDolar(value)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)

}

func QuatationHandler(res http.ResponseWriter, req *http.Request) {
	api := integration.NewExchangeRateApi("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	service := service.NewQuotationService(*api)
	result, err := service.GetCurrentExchange()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/convert-high", ConvertHighHandler)
	mux.HandleFunc("/cotacao", QuatationHandler)

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Server is listening on port 8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
