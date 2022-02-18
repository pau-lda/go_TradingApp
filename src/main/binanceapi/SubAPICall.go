package binanceapi

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

//TODO implement all other market data endpoints
