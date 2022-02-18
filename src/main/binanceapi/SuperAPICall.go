package binanceapi

import (
	"encoding/json"
	"fmt"
	"go_TradingApp/main/binanceapi/constants"
	"go_TradingApp/main/binanceapi/enums"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type superAPICall struct{}
type privateSuperAPICall struct{}

func finalRequest[E any](reqtype enums.RequestType, url *url.URL) (ret E) {

	client := &http.Client{}

	req, err := http.NewRequest(strings.ToUpper(string(reqtype)), url.String(), nil)

	//check if call is a valid format
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	res, err := client.Do(req)

	//check if call was successfully executed
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseBody, err := ioutil.ReadAll(res.Body)

	//check if data from call was successfully read
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(responseBody, &ret)

	//check if conversion from json to object was successfully executed
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return

}

func makeGETRequest[E any](url *url.URL) (ret E) { //accepts any <E> type and also returns that type

	url.Scheme = constants.HTTPS
	url.Host = constants.BINANCE_TEST_NET

	return finalRequest[E](enums.GET, url)

}
