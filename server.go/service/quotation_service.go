package service

import (
	integration "github.com/RomeroGabriel/server.go/integration"
)

type QuotationService struct {
	api integration.ExchangeRateApi
}

func NewQuotationService(api integration.ExchangeRateApi) *QuotationService {
	return &QuotationService{
		api: api,
	}
}

func (service QuotationService) CalculateRealToDolar(real_value float64) (result float64, err error) {
	data, err := service.api.GetRealDolarRate()
	if err != nil {
		return 0.0, err
	}
	return real_value * data.High, nil
}
