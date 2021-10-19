package kraken

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/rudyjcruz831/backendCrypto/utils"
)

type ResponsePriceBTC struct {
	Error  []string `json:"erorr"`
	Result Btc      `json:"result"`
}

type ResponsePriceETH struct {
	Error  []string `json:"erorr"`
	Result Eth      `json:"result"`
}

type Btc struct {
	Exchange Exchange `json:"XXBTZUSD"`
}

type Eth struct {
	Exchange Exchange `json:"XETHZUSD"`
}

type Exchange struct {
	A []string `json:"a"`
	B []string `json:"b"`
	C []string `json:"c"`
	V []string `json:"v"`
	P []string `json:"p"`
	T []string `json:"t"`
	L []string `json:"l"`
	H []string `json:"h"`
	O string   `json:"o"`
}

// get the info for USD to BTC
// https://api.kraken.com/0/public/Ticker?pair=XBTUSD

// get the info for USD to ETH
// "https://api.kraken.com/0/public/Ticker?pair=XETHZUSD"
// {
// 	"error":[],
// 	"result":{
// 		"XETHZUSD":{
// 			"a":["3884.88000","1","1.000"],
// 			"b":["3884.87000","7","7.000"], // Sell price
// 			"c":["3884.93000","0.14754333"],
// 			"v":["1909.39463462","23895.66354427"],
// 			"p":["3870.36079","3798.31333"],
// 			"t":[1529,21504],
// 			"l":["3842.00000","3640.48000"],
// 			"h":["3885.45000","3919.92000"],
// 			"o":"3846.90000"} // Buy price
// 		}
// }
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

func GetBuyandSellPriceBTCtoUSD() (string, string) {
	req := newRequest("api.kraken.com/0/public/Ticker?pair=XBTUSD")

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

	var result ResponsePriceBTC
	if err := json.Unmarshal(data, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON GetBuyandSellPriceBTCtoUSD")
	}

	buyPrice := result.Result.Exchange.O
	sellPrice := result.Result.Exchange.C[0]

	buyPriceFloat, err := strconv.ParseFloat(buyPrice, 64)
	if err != nil {
		log.Fatalf("%e\n", err)
	}
	sellPriceFloat, err := strconv.ParseFloat(sellPrice, 64)
	if err != nil {
		log.Fatalf("%e\n", err)
	}

	return fmt.Sprintf("%.2f", buyPriceFloat), fmt.Sprintf("%.2f", sellPriceFloat)
}

func GetBuyAndSellPriceETHtoUSD() (string, string) {
	req := newRequest("api.kraken.com/0/public/Ticker?pair=XETHZUSD")

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

	var result ResponsePriceETH
	if err := json.Unmarshal(data, &result); err != nil { // Parse []byte to go struct pointer
		fmt.Println("Can not unmarshal JSON GetBuyAndSellPriceETHtoUSD")
	}
	// fmt.Println(string(data))
	// fmt.Println(result)

	buyPrice := result.Result.Exchange.O
	sellPrice := result.Result.Exchange.C[0]

	buyPriceFloat, err := strconv.ParseFloat(buyPrice, 64)
	if err != nil {
		log.Fatalf("%e\n", err)
	}
	sellPriceFloat, err := strconv.ParseFloat(sellPrice, 64)
	if err != nil {
		log.Fatalf("%e\n", err)
	}
	return fmt.Sprintf("%.2f", buyPriceFloat), fmt.Sprintf("%.2f", sellPriceFloat)
}
