package main

import (
	"fmt"
	"go_TradingApp/main/binanceapi"
	"testing"
)

func TestCall(t *testing.T) {

	var test binanceapi.SubAPICall

	time := test.GetServerTime()
	fmt.Println(time.GetServerTime())
}
