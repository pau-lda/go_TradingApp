package binanceapi

import (
	"encoding/json"
	"fmt"
	calls "go_TradingApp/main/binanceapi/marketDataEndpoints"
	"io/ioutil"
	"net/http"
	"os"
	_ "os"
)

type SuperAPICall struct{}

func (SuperAPICall) GetServerTime() calls.ServerTime {
	var servertime calls.ServerTime

	response, err := http.Get("https://binance.com/api/v3/time")

	//check if call was successfully executed
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseBody, err := ioutil.ReadAll(response.Body)

	//check if data from call was successfully read
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(responseBody, &servertime)

	//check if conversion from json to object was successfully executed
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return servertime
}
