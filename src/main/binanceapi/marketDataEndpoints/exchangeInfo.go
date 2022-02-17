package marketDataEndpoints

type ExchangeInfo struct {
	Timezone        string        `json:"timezone"`
	ServerTime      int64         `json:"serverTime"`
	RateLimits      []interface{} `json:"rateLimits"`
	ExchangeFilters []string      `json:"exchangeFilters"`
	Symbols         []Symbol      `json:"symbols"`
}

func (e *ExchangeInfo) GetTimezone() string {
	return e.Timezone
}

func (e *ExchangeInfo) GetServerTime() int64 {
	return e.ServerTime
}

func (e *ExchangeInfo) GetExchangeFilters() []string {
	return e.ExchangeFilters
}

func (e *ExchangeInfo) GetSymbols() []Symbol {
	return e.Symbols
}
