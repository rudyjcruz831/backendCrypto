# backendCrypto
    - This backend uses different APIs to fetch data for Bitcoin and Ethereum
    
    Kraken 
    [a link](https://docs.kraken.com/rest/)
    
    apis used 
    GET https://api.kraken.com/0/public/Ticker?pair=XBTUSD
    GET https://api.kraken.com/0/public/Ticker?pair=XETHZUSD

    
    Coibase API
    [a link](https://developers.coinbase.com/api/v2)

    apis used 
    GET https://api.coinbase.com/v2/prices/:currency_pair/buy
    GET https://api.coinbase.com/v2/prices/:currency_pair/sell


# Installations

To run backend, you need to install Go and set your Go workspace first.

1. The first need Go installed (version 1.14+ is required)

# Run
    $ go mod tidy
    $ go run main.go