package integration

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type ExchangeRateApi struct {
	url_api string
}

func NewExchangeRateApi(url_api string) *ExchangeRateApi {
	return &ExchangeRateApi{
		url_api: url_api,
	}
}

func callGetApi(api ExchangeRateApi) (*ApiResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", api.url_api, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("Request Timeout consulting the " + api.url_api + " API")
		return nil, context.Canceled
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var result ApiResponse
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	select {
	case <-ctx.Done():
		log.Println("Request Timeout consulting the " + api.url_api + " API")
		return nil, context.Canceled
	default:
		return &result, nil
	}
}

func (api ExchangeRateApi) GetRealDolarRate() (*ExchangeRateResult, error) {
	result, err := callGetApi(api)
	if err != nil {
		return nil, err
	}
	exchangeRateResult := ParserApiResponseToExchangeRateResult(*result)
	return &exchangeRateResult, nil
}

func (api ExchangeRateApi) GetExchangeValue() (float64, error) {
	result, err := callGetApi(api)
	if err != nil {
		return 0.0, err
	}
	exchangeValue := convertStringToFloat(result.DataResponse.Bid)
	return exchangeValue, nil
}
