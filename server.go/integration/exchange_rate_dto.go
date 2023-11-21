package integration

import (
	"fmt"
	"strconv"
)

type ApiResponse struct {
	DataResponse struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

type ExchangeRateResult struct {
	Code   string  `json:"code"`
	Codein string  `json:"codein"`
	Name   string  `json:"name"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
}

func ParserApiResponseToExchangeRateResult(data ApiResponse) ExchangeRateResult {
	exchangeRateResult := ExchangeRateResult{
		Code:   data.DataResponse.Code,
		Codein: data.DataResponse.Codein,
		Name:   data.DataResponse.Name,
		High:   convertStringToFloat(data.DataResponse.High),
		Low:    convertStringToFloat(data.DataResponse.Low),
	}
	return exchangeRateResult
}

func convertStringToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		// handle error
		fmt.Println(err)
		return 0
	}
	return f
}
