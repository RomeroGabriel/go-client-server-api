package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/mattn/go-sqlite3"

	integration "github.com/RomeroGabriel/server.go/integration"
	repository "github.com/RomeroGabriel/server.go/repository"
	service "github.com/RomeroGabriel/server.go/service"
)

func createDataBase() (*sql.DB, error) {
	fileName := "exchange.db"
	finalFileName := "./" + fileName

	path, err := os.Getwd()
	if err == nil {
		finalFileName = path + "/" + fileName
	} else {
		log.Print(err.Error())
	}
	return sql.Open("sqlite3", finalFileName)
}

func ConvertHighHandler(res http.ResponseWriter, req *http.Request) {
	api := integration.NewExchangeRateApi("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	db, err := createDataBase()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	repo := repository.NewExhangeRateRepository(db)
	service := service.NewQuotationService(*api, *repo)

	queryData := req.URL.Query()
	if !queryData.Has("value") {
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	value_par := req.URL.Query().Get("value")
	value, err := strconv.ParseFloat(value_par, 64)
	if err != nil {
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
	defer db.Close()
}

func QuatationHandler(res http.ResponseWriter, req *http.Request) {
	api := integration.NewExchangeRateApi("https://economia.awesomeapi.com.br/json/last/USD-BRL")
	db, err := createDataBase()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	repo := repository.NewExhangeRateRepository(db)
	service := service.NewQuotationService(*api, *repo)

	result, err := service.GetCurrentExchange()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(err.Error()))
		return
	}
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(result)
	defer db.Close()
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
