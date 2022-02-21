package binanceapi

/**
Class only gets called by SubAPICall
Never call methods of SuperAPICall directly
*/

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

//requires:
//	return type specification and enum of RequestType as well as the url to call
//structure:
//
//	finalRequest[YOURTYPE](enum YOURREQUESTTYPE, url YOURURL)
//
//returns:
//	struct with response of type <E>
func finalRequest[E any](reqtype enums.RequestType, url *url.URL) (ret E) {

	client := &http.Client{}

	req, err := http.NewRequest(strings.ToUpper(string(reqtype)), url.String(), nil)

	//check if call is a valid format
	if err != nil {
		fmt.Println("No valid url")
		fmt.Print(err.Error())
		os.Exit(1)
	}

	res, err := client.Do(req)

	//check if call was successfully executed
	if err != nil {
		fmt.Println("Call was not successful")
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseBody, err := ioutil.ReadAll(res.Body)

	//check if data from call was successfully read
	if err != nil {
		fmt.Println("Call was not successfully read")
		fmt.Print(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(responseBody, &ret)

	//check if conversion from json to object was successfully executed
	if err != nil {
		fmt.Println("JSON parser was not succesful")
		fmt.Print(err.Error())
		os.Exit(1)
	}

	return //returns ret of type E

}

//requires:
//	return type specification as well as the path with parameters if the url to call
//combines given path and parameters with scheme and host and call finaRequest
//structure:
//
//	makeGETRequest[YOURTYPE](url YOURURL)
//
//returns:
//	struct with GET response of type <E>
func makeGETRequest[E any](url *url.URL) (ret E) {
	url.Scheme = constants.HTTPS
	url.Host = constants.BINANCE_TEST_NET
	return finalRequest[E](enums.GET, url)

}

//requires:
//	return type specification as well as the path with parameters if the url to call
//combines given path and parameters with scheme and host and call finaRequest
//structure:
//
//	makeGETRequest[YOURTYPE](url YOURURL)
//
//returns:
//	struct with POST response of type <E>
func makePOSTRequest[E any](url *url.URL) (ret E) {
	url.Scheme = constants.HTTPS
	url.Host = constants.BINANCE_TEST_NET
	return finalRequest[E](enums.POST, url)
}
