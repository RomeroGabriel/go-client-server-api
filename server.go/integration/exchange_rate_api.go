package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type ExchangeRateApi struct {
	url_api string
}

func NewExchangeRateApi(url_api string) *ExchangeRateApi {
	return &ExchangeRateApi{
		url_api: url_api,
	}
}

func (api ExchangeRateApi) GetRealDolarRate() (*ExchangeRateResult, error) {
	req, err := http.Get(api.url_api)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	res, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	var result ApiResponse
	err = json.Unmarshal(res, &result)
	if err != nil {
		return nil, err
	}
	exchangeRateResult := ExchangeRateResult{
		Code:   result.DataResponse.Code,
		Codein: result.DataResponse.Codein,
		Name:   result.DataResponse.Name,
		High:   convertStringToFloat(result.DataResponse.High),
		Low:    convertStringToFloat(result.DataResponse.Low),
	}
	return &exchangeRateResult, nil
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
