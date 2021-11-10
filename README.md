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


# API 
Routes for API examples
## Get list of sell and buy price for ETH and BTC
### Request
GET /info
    curl http://localhost:8081/info
### Response
    HTTP/1.1 200 OK
    Date: Thu, 24 Feb 2011 12:36:30 GMT
    Status: 200 OK
    Connection: close
    Content-Type: application/json
    {
        "Exchanges": [
            {
                "source": "Coinbase",
                "name": "Bitcoin",
                "buy": "67451.12",
                "sell": "66775.75"
            },
            {
                "source": "Coinbase",
                "name": "Ethereum",
                "buy": "4754.21",
                "sell": "4706.33"
            },
            {
                "source": "Kraken",
                "name": "Bitcoin",
                "buy": "67105.00",
                "sell": "66497.69"
            },
            {
                "source": "Kraken",
                "name": "Ethereum",
                "buy": "4730.46",
                "sell": "4687.89"
            }
        ]
    }
    


## Get the best prices from Kraken and Coinbase
### Request 
GET /best
    curl http://localhost:8081/best
### Response
    HTTP/1.1 200 OK
    Status: 200 OK
    Connection: close
    Content-Type: application/json

    {
        "best_prices": [
            {
                "name": "Ethereum",
                "best_sell": "4702.93",
                "sell_source": "Coinbase",
                "best_buy": "4726.81",
                "buy_source": "Kraken"
            },
            {
                "name": "Bitcoin",
                "best_sell": "66771.08",
                "sell_source": "Coinbase",
                "best_buy": "67094.70",
                "buy_source": "Kraken"
            }
        ]
    }
    


    
# Questionnaire:
    1. One of the biggest things was design it had to be limieted for example I did not add security. 
    Making test for backend or frontend. Also request for app was simple did not really need to make a robust web.
    
    2. I do not thing I over designed. I did worked more on the designed of the backend for the reason that I understand how to better created it.
    
    3. If I had to scale for 100 users I would have to upgarade the heroku servers to handle more traffic. Second I would think of a 
    catching system or maybe some kind of load balancer. The last thing is better design for scaling backend and frontend. 
    
    4. I needed more time to understand frontend design and better practices to develop a much more robust frontend. 
    For example frontend I would like more time to understand hooks like useState(). With backend I need more time to create proper testing for 
    endpoints and add a layer of security. For example for backend I am calling external api twice where I might be able to just call external APIs once.
