package main

import (
	"fmt"
	"go_TradingApp/main/binanceapi"
	"testing"
)

func TestCall(t *testing.T) {

	var test binanceapi.SubAPICall

	time := test.GetServerTime()
	info := test.GetExchangeInfo()
	fmt.Println(time.GetServerTime())
	for _, element := range info.Symbols { //iterate over symbol array
		fmt.Println(element.Symbol)
	}

}
