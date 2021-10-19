package coinbase

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/rudyjcruz831/backendCrypto/utils"
)

type ResponsePrice struct {
	Data Price `json:"data"`
}

type Price struct {
	Base     string `json:"base"`
	Currency string `json:"currency"`
	Amount   string `json:"amount"`
}

func newRequest(endpoint string) *http.Request {
	env := utils.GetENV{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s", endpoint), nil)
	if err != nil {
		log.Fatalf("error %e", err)
	}
	// fmt.Println(req)
	apiKey := env.GoDotEnvVariable("COINBASE_KEY")
	secret := env.GoDotEnvVariable("COINBASE_SECRET")
	req.Header.Add(apiKey, secret)
	return req
}

// get buy prices for BTC to USD
func GetBuyPriceBTCtoUSD() string {

	req := newRequest("api.coinbase.com/v2/prices/BTC-USD/buy")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(res.StatusCode)
	data, errBody := ioutil.ReadAll(res.Body)
	if errBody != nil {
		log.Fatal(errBody)
	}
	res.Body.Close()

	var result ResponsePrice
	if err := json.Unmarshal(data, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON GetBuyPriceBTCtoUSD")
	}
	// fmt.Println(result.Data)
	return result.Data.Amount
}

// get buy prices for ETH to USD
func GetBuyPriceETHtoUSD() string {

	req := newRequest("api.coinbase.com/v2/prices/ETH-USD/buy")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, errBody := ioutil.ReadAll(res.Body)
	if errBody != nil {
		log.Fatal(errBody)
	}
	res.Body.Close()

	var result ResponsePrice
	if err := json.Unmarshal(data, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON GetBuyPriceETHtoUSD")
	}
	// fmt.Println(result.Data)
	return result.Data.Amount
}

func GetSellPriceBTCtoUSD() string {

	req := newRequest("api.coinbase.com/v2/prices/BTC-USD/sell")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, errBody := ioutil.ReadAll(res.Body)
	if errBody != nil {
		log.Fatal(errBody)
	}
	res.Body.Close()

	var result ResponsePrice
	if err := json.Unmarshal(data, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON GetSellPriceBTCtoUSD")
	}
	// fmt.Println(result.Data)
	return result.Data.Amount
}

func GetSellPriceETHtoUSD() string {

	req := newRequest("api.coinbase.com/v2/prices/ETH-USD/sell")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	data, errBody := ioutil.ReadAll(res.Body)
	if errBody != nil {
		log.Fatal(errBody)
	}
	res.Body.Close()

	var result ResponsePrice
	if err := json.Unmarshal(data, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON GetSellPriceETHtoUSD")
	}
	// fmt.Println(result.Data)
	return result.Data.Amount
}
