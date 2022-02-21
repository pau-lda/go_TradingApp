package binanceapi

/**
Class to call when making a request
*/

import (
	calls "go_TradingApp/main/binanceapi/marketDataEndpoints"
	"net/url"
)

type SubAPICall struct {
}

func (*SubAPICall) GetServerTime() calls.ServerTime {
	var url url.URL
	url.Path = "/api/v3/time"
	time := makeGETRequest[calls.ServerTime](&url)

	//fmt.Println(time)
	return time
}

func (*SubAPICall) GetExchangeInfo() calls.ExchangeInfo {
	url := url.URL{Path: "/api/v3/exchangeInfo"}

	info := makeGETRequest[calls.ExchangeInfo](&url)
	return info
}

//TODO implement all other market data endpoints
