package main

import (
	"encoding/json"
	"fmt"
	marketData "go_TradingApp/main/binanceapi/marketDataEndpoints"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"testing"
)

func TestCall(t *testing.T) {

	response, err := http.Get("https://binance.com/api/v3/time")
	response2, err2 := http.Get("https://binance.com/api/v3/exchangeInfo")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	if err2 != nil {
		fmt.Print(err2.Error())
		os.Exit(2)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	responseData2, err2 := ioutil.ReadAll(response2.Body)
	if err2 != nil {
		log.Fatal(err2)
	}

	//fmt.Println(string(responseData2))

	var ResponseObject marketData.ServerTime
	var ResponseObject2 marketData.ExchangeInfo

	json.Unmarshal(responseData, &ResponseObject)
	json.Unmarshal(responseData2, &ResponseObject2)

	fmt.Println(ResponseObject.GetServerTimeFormatted())
	_symbols := ResponseObject2.GetSymbols()
	for i := 0; i < len(_symbols); i++ {
		fmt.Println(_symbols[i].Symbol)
	}
}
