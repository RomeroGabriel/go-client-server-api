package service

import (
	integration "github.com/RomeroGabriel/server.go/integration"
	repository "github.com/RomeroGabriel/server.go/repository"
)

type QuotationService struct {
	api        integration.ExchangeRateApi
	repository repository.ExhangeRateRepository
}

func NewQuotationService(api integration.ExchangeRateApi, repository repository.ExhangeRateRepository) *QuotationService {
	return &QuotationService{
		api:        api,
		repository: repository,
	}
}

func (service QuotationService) CalculateRealToDolar(real_value float64) (result float64, err error) {
	data, err := service.api.GetRealDolarRate()
	if err != nil {
		return 0.0, err
	}
	return real_value * data.High, nil
}

func (service QuotationService) GetCurrentExchange() (float64, error) {
	data, err := service.api.GetExchangeValue()
	if err != nil {
		return 0.0, err
	}

	err = service.repository.CreateExchange(data)
	if err != nil {
		return 0.0, err
	}

	return data, nil
}
