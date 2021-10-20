package kraken

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// Response type Different for mapping BTC
type ResponsePriceBTC struct {
	Error  []string `json:"erorr"`
	Result Btc      `json:"result"`
}

// Response type Different for mapping ETH
type ResponsePriceETH struct {
	Error  []string `json:"erorr"`
	Result Eth      `json:"result"`
}

// struct to map to BTC b/c different for each exchange
type Btc struct {
	Exchange Exchange `json:"XXBTZUSD"`
}

// struct to map to ETH b/c different for each exchange
type Eth struct {
	Exchange Exchange `json:"XETHZUSD"`
}

// struct to map differnt Prices for the same for all Currencies
type Exchange struct {
	A []string `json:"a"`
	B []string `json:"b"`
	C []string `json:"c"`
	V []string `json:"v"`
	P []string `json:"p"`
	T []int    `json:"t"`
	L []string `json:"l"`
	H []string `json:"h"`
	O string   `json:"o"`
}

// routes to kraken API to get price for ETH and BTC
// https://api.kraken.com/0/public/Ticker?pair=XBTUSD
// https://api.kraken.com/0/public/Ticker?pair=XETHZUSD

// Sample of what kraken API returns
// {
// 	"error":[],
// 	"result":{
// 		"XETHZUSD":{
// 			"a":["3884.88000","1","1.000"], string
// 			"b":["3884.87000","7","7.000"], // Sell price
// 			"c":["3884.93000","0.14754333"], string
// 			"v":["1909.39463462","23895.66354427"], string
// 			"p":["3870.36079","3798.31333"], string
// 			"t":[1529,21504], int
// 			"l":["3842.00000","3640.48000"], string
// 			"h":["3885.45000","3919.92000"], string
// 			"o":"3846.90000"} // Buy price
// 		}
// }

// create the request to make http call to server takes in
// name of endpoint after https://
func newRequest(endpoint string) *http.Request {
	// env := utils.GetENV{}

	req, err := http.NewRequest("GET", fmt.Sprintf("https://%s", endpoint), nil)
	if err != nil {
		log.Fatalf("error %e", err)
	}
	// fmt.Println(req)
	// apiKey := env.GoDotEnvVariable("COINBASE_KEY")
	// secret := env.GoDotEnvVariable("COINBASE_SECRET")
	// req.Header.Add(apiKey, secret)
	return req
}

// gets the json for price of BTC to USD from kraken
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
		fmt.Printf("Can not unmarshal JSON GetBuyandSellPriceBTCtoUSD: %s\n", err)
	}

	buyPrice := result.Result.Exchange.A[0]
	sellPrice := result.Result.Exchange.C[0]

	f, err := strconv.ParseFloat(sellPrice, 64)
	takenOff := f * float64(0.009)

	buyPriceFloat, err := strconv.ParseFloat(buyPrice, 64)
	if err != nil {
		log.Fatalf("%e\n", err)
	}
	sellPriceFloat, err := strconv.ParseFloat(sellPrice, 64)
	if err != nil {
		log.Fatalf("%e\n", err)
	}
	sellPriceFloat -= takenOff
	return fmt.Sprintf("%.2f", buyPriceFloat), fmt.Sprintf("%.2f", sellPriceFloat)
}

// gets the json for price of ETH to USD from kraken
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
		fmt.Printf("Can not unmarshal JSON GetBuyAndSellPriceETHtoUSD: %s\n", err)
	}
	// fmt.Println(string(data))
	// fmt.Println(result)

	buyPrice := result.Result.Exchange.A[0]
	sellPrice := result.Result.Exchange.C[0]
	// krakenFee := 0.009
	f, err := strconv.ParseFloat(sellPrice, 64)
	takenOff := f * float64(0.009)
	// s := fmt.Sprintf("%f", takenOff) // s == "123.456000"
	// fmt.Println(s)

	buyPriceFloat, err := strconv.ParseFloat(buyPrice, 64)
	if err != nil {
		log.Fatalf("%e\n", err)
	}
	sellPriceFloat, err := strconv.ParseFloat(sellPrice, 64)
	if err != nil {
		log.Fatalf("%e\n", err)
	}
	sellPriceFloat -= takenOff
	return fmt.Sprintf("%.2f", buyPriceFloat), fmt.Sprintf("%.2f", sellPriceFloat)
}
