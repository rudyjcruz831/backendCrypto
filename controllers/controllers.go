package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/backendCrypto/coinbase"
	"github.com/rudyjcruz831/backendCrypto/kraken"
)

type Exchanges struct {
	Crypto []Crypto `json:"Exchanges"`
}

type Crypto struct {
	Source string `json:"source"`
	Name   string `json:"name"`
	Buy    string `json:"buy"`
	Sell   string `json:"sell"`
}

type BestExchanges struct {
	BestPrices []BestPrices `json:"best_prices"`
}

type BestPrices struct {
	Name       string `json:"name"`
	BestSell   string `json:"best_sell"`
	SellSource string `json:"sell_source"`
	BestBuy    string `json:"best_buy"`
	BuySource  string `json:"buy_source"`
}

// GetJSON make a get request and returns
// prices for ETH and BTC from Kraken and Coinbase API
func GetCryptoInfo(c *gin.Context) {

	// functions that handle getting prices from coinbase
	buyPriceBTCcoinbase := coinbase.GetBuyPriceBTCtoUSD()
	sellPriceBTCcoinbase := coinbase.GetSellPriceBTCtoUSD()
	buyPriceETHcoinbase := coinbase.GetBuyPriceETHtoUSD()
	sellPriceETHcoinbase := coinbase.GetSellPriceETHtoUSD()

	// fuctions that handle getting prices from kraken
	buyPriceBTCkraken, sellPriceBTCkraken := kraken.GetBuyandSellPriceBTCtoUSD()
	buyPriceETHkraken, sellPriceETHkraken := kraken.GetBuyAndSellPriceETHtoUSD()

	bitcoinCoinbase := Crypto{
		Name:   "Bitcoin",
		Source: "Coinbase",
		Buy:    buyPriceBTCcoinbase,
		Sell:   sellPriceBTCcoinbase,
	}
	ethereumCoinbase := Crypto{
		Name:   "Ethereum",
		Source: "Coinbase",
		Buy:    buyPriceETHcoinbase,
		Sell:   sellPriceETHcoinbase,
	}
	bitcoinKraken := Crypto{
		Name:   "Bitcoin",
		Source: "Kraken",
		Buy:    buyPriceBTCkraken,
		Sell:   sellPriceBTCkraken,
	}
	ethereumKraken := Crypto{
		Name:   "Ethereum",
		Source: "Kraken",
		Buy:    buyPriceETHkraken,
		Sell:   sellPriceETHkraken,
	}

	ex := Exchanges{}
	ex.Crypto = append(ex.Crypto, bitcoinCoinbase)
	ex.Crypto = append(ex.Crypto, ethereumCoinbase)
	ex.Crypto = append(ex.Crypto, bitcoinKraken)
	ex.Crypto = append(ex.Crypto, ethereumKraken)

	c.IndentedJSON(http.StatusCreated, ex)
}

// GetJSON make a get request and returns best prices
// for ETH and BTC from Kraken and Coinbase API
func GetBestPrices(c *gin.Context) {
	buyPriceBTCcoinbase := coinbase.GetBuyPriceBTCtoUSD()

	sellPriceBTCcoinbase := coinbase.GetSellPriceBTCtoUSD()

	buyPriceETHcoinbase := coinbase.GetBuyPriceETHtoUSD()
	sellPriceETHcoinbase := coinbase.GetSellPriceETHtoUSD()

	buyPriceBTCkraken, sellPriceBTCkraken := kraken.GetBuyandSellPriceBTCtoUSD()
	buyPriceETHkraken, sellPriceETHkraken := kraken.GetBuyAndSellPriceETHtoUSD()

	bitcoinCoinbase := Crypto{
		Name:   "Bitcoin",
		Source: "Coinbase",
		Buy:    buyPriceBTCcoinbase,
		Sell:   sellPriceBTCcoinbase,
	}
	ethereumCoinbase := Crypto{
		Name:   "Ethereum",
		Source: "Coinbase",
		Buy:    buyPriceETHcoinbase,
		Sell:   sellPriceETHcoinbase,
	}
	bitcoinKraken := Crypto{
		Name:   "Bitcoin",
		Source: "Kraken",
		Buy:    buyPriceBTCkraken,
		Sell:   sellPriceBTCkraken,
	}
	ethereumKraken := Crypto{
		Name:   "Ethereum",
		Source: "Kraken",
		Buy:    buyPriceETHkraken,
		Sell:   sellPriceETHkraken,
	}

	eCoinbaseBuy, _ := strconv.ParseFloat(ethereumCoinbase.Buy, 64)
	eKrakenBuy, _ := strconv.ParseFloat(ethereumKraken.Buy, 64)
	eCoinbaseSell, _ := strconv.ParseFloat(ethereumCoinbase.Sell, 64)
	eKrakenSell, _ := strconv.ParseFloat(ethereumKraken.Sell, 64)

	bKrakenBuy, _ := strconv.ParseFloat(bitcoinKraken.Buy, 64)
	bCoinbaseBuy, _ := strconv.ParseFloat(bitcoinCoinbase.Buy, 64)
	bKrakenSell, _ := strconv.ParseFloat(bitcoinKraken.Sell, 64)
	bCoinbaseSell, _ := strconv.ParseFloat(bitcoinCoinbase.Sell, 64)

	exB := BestExchanges{}
	eBest := BestPrices{Name: "Ethereum"}
	bBest := BestPrices{Name: "Bitcoin"}
	if eCoinbaseBuy < eKrakenBuy {
		eBest.BestBuy = ethereumCoinbase.Buy
		eBest.BuySource = "Coinbase"
	} else {
		eBest.BestBuy = ethereumKraken.Buy
		eBest.BuySource = "Kraken"
	}

	if eCoinbaseSell > eKrakenSell {
		eBest.BestSell = ethereumCoinbase.Sell
		eBest.SellSource = "Coinbase"
	} else {
		eBest.BestSell = ethereumKraken.Sell
		eBest.SellSource = "Kraken"
	}

	if bCoinbaseBuy < bKrakenBuy {
		bBest.BestBuy = bitcoinCoinbase.Buy
		bBest.BuySource = "Coinbase"
	} else {
		bBest.BestBuy = bitcoinKraken.Buy
		bBest.BuySource = "Kraken"
	}

	if bCoinbaseSell > bKrakenSell {
		bBest.BestSell = bitcoinCoinbase.Sell
		bBest.SellSource = "Coinbase"
	} else {
		bBest.BestSell = bitcoinKraken.Sell
		bBest.SellSource = "Kraken"
	}

	exB.BestPrices = append(exB.BestPrices, eBest)
	exB.BestPrices = append(exB.BestPrices, bBest)

	c.IndentedJSON(http.StatusCreated, exB)
}

func GetSomething(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, nil)
}
