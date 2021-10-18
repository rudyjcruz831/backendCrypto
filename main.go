package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rudyjcruz831/backendCrypto/coinbase"
	"github.com/rudyjcruz831/backendCrypto/kraken"
)

func main() {

	router := gin.Default()
	router.GET("/info", getCryptoInfo)

	router.Run("localhost:8080")

}

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
func getCryptoInfo(c *gin.Context) {
	// var newAlbum album

	buyPriceBTCcoinbase := coinbase.GetBuyPriceBTCtoUSD()

	sellPriceBTCcoinbase := coinbase.GetSellPriceBTCtoUSD()

	buyPriceETHcoinbase := coinbase.GetBuyPriceETHtoUSD()
	sellPriceETHcoinbase := coinbase.GetSellPriceETHtoUSD()

	buyPriceBTCkraken, sellPriceBTCkraken := kraken.GetBuyandSellPriceBTCtoUSD()
	buyPriceETHkraken, sellPriceETHkraken := kraken.GetBuyAndSellPriceETHtoUSD()

	bitcoinCoinbase := Crypto{
		Name:   "Bitcoin",
		Source: "CoinBase",
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

	// // Call BindJSON to bind the received JSON to
	// // newAlbum.
	// if err := c.BindJSON(&newAlbum); err != nil {
	// 	return
	// }

	// // Add the new album to the slice.
	// albums = append(albums, newAlbum)
	// c.IndentedJSON(http.StatusCreated, newAlbum)

}
