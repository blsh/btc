## Xuy-BTC


### About
Thi is an atempt to write a software which will parse diferent Bitcoin trading
data, seriliaze it and feed it to other services like neuronal networks,
databases and trading bots.

### Goals
#### Goal 1
1. [ ] Parse bitcoincharts.com trading data from telnet interface
    1. [X] Log all incoming raw data
    2. [ ] Log all USD data in a format which can be parsed by a simple Neuroph
       MLP network with 4 Inputs (symbol, price, value, timestamp) and 3 outputs
       (priceGoesUp, priceWentDown, priceSame)
        1. [ ] Normalize data. All values should be in range of 0 - 1
        2. [ ] Write data to a csv file

### NOTE!
#### Currencies
Everytime we talk about price changes, we do it in some unspecified currency
(CUR).  The whole code is completly currency agnostic, so we can use it for USD,
EUR, RUR, $WHATVER
