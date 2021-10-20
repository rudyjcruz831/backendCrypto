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
    
 # Questionnaire:
    1. One of the biggest things was design it had to be limieted for example I did not add security. Making test for backend ot frontend.
    
    2. I do not thing I over designed. I did worked more on the designed of the backend for the reason that I understand how to better created it.
    
    3. If I had to scale for 100 users I would have to upgarade the heroku servers to handle more traffic. Second I would think of a catching system or maybe load 
    balancing. The last thing is better design for scaling backend and frontend. 
    
    4. 
