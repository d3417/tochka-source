package marketplace

import (
	"github.com/d3417/tochka-source/modules/apis"
	"sync"
)

/*
	Models
*/

var (
	CRYPTO_CURRENCIES = []string{"ETH", "BTC"}
	FIAT_CURRENCIES   = []string{"USD", "EUR", "AUD", "GBP", "RUB"}
	CURRENCY_RATES    = sync.Map{} // map[string]map[string]float64{}
)

func UpdateCurrencyRates() {
	for _, currency := range CRYPTO_CURRENCIES {
		rates, err := apis.GetCurrencyRates(currency)
		if err != nil {
			continue
		}
		CURRENCY_RATES.Store(currency, rates)
		// CURRENCY_RATES[currency] = rates
	}

	for _, currency := range FIAT_CURRENCIES {
		rates, err := apis.GetCurrencyRates(currency)
		if err != nil {
			continue
		}
		// CURRENCY_RATES[currency] = rates
		CURRENCY_RATES.Store(currency, rates)
	}
}

func GetCurrencyRate(baseCurrency, targetCurrency string) float64 {
	if _rates, ok1 := CURRENCY_RATES.Load(baseCurrency); ok1 {
		rates := _rates.(map[string]float64)
		if rate, ok2 := rates[targetCurrency]; ok2 {
			return rate
		}
	}
	return 1.0
}
