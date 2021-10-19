package controllers

import (
	"net/http"

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

// postAlbums adds an album from JSON received in the request body.
func GetCryptoInfo(c *gin.Context) {
	// var newAlbum album

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

	ex := Exchanges{}
	ex.Crypto = append(ex.Crypto, bitcoinCoinbase)
	ex.Crypto = append(ex.Crypto, ethereumCoinbase)
	ex.Crypto = append(ex.Crypto, bitcoinKraken)
	ex.Crypto = append(ex.Crypto, ethereumKraken)

	c.IndentedJSON(http.StatusCreated, ex)
}
