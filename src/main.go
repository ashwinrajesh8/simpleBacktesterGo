package main

import (
	"net/http"
	"fmt"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
)

type results struct {
	stock, strat float64
}

type group struct {
	indicator, base float64
}

type trigger struct {
	isAbove, isBelow, crossesAbove, crossesBelow bool
	fiftyDayMA, twohundredDayMA bool						// offered indicators
	appendConstant float64
	multiplier float64
}

type indicator struct{
	fiftyDayMA, twohundredDayMA bool
}


  type AggregateJSON struct {
	Ticker       string `json:"ticker"`
	QueryCount   int    `json:"queryCount"`
	ResultsCount int    `json:"resultsCount"`
	Adjusted     bool   `json:"adjusted"`
	Results      []struct {
		V  int     `json:"v"`
		Vw float64 `json:"vw"`
		O  float64 `json:"o"`
		C  float64 `json:"c"`
		H  float64 `json:"h"`
		L  float64 `json:"l"`
		T  int64   `json:"t"`
		N  int     `json:"n"`
	} `json:"results"`
	Status    string `json:"status"`
	RequestID string `json:"request_id"`
	Count     int    `json:"count"`
}

var apiKey = ""		// insert Polygon.io API key (up to 100 req/day)

var dataPoints int = 0						// used to keep track of how many plots on graph
var performance []results					// appended by strat runner containing benchmark of stock and porfolio, used to graph

var buyTrigger []trigger
var sellTrigger []trigger

var ticker string
var interval int

var timey = time.Now()													// used as ref for indicator funcs
var Timestamp = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)						// time rn
var TimeIncrementor = strconv.FormatInt(time.Now().AddDate(0,0,-8).UTC().UnixNano(), 10)		// go back 8 days (backtesting period)

func postBacktest(c *gin.Context) {
	ticker = c.Param("ticker")
	tempInterval, erry := strconv.Atoi(c.Param("interval"))
	if erry == nil {
        fmt.Printf("Request error.")
    }
	interval = tempInterval
	fmt.Println("got here")
	Timestamp = Timestamp[:len(Timestamp)-6]											// fo some reason, removin last 6 digits gives correct unix time?!
	TimeIncrementor = TimeIncrementor[:len(TimeIncrementor)-6]									// fo some reason, removin last 6 digits gives correct unix time?!

	// benchmarkStrat - stores json struct object for first metric (for however many intervals fit in week), to be run against, need to identify buy/sell, crosses/is above/below

	netStratReturn := fiftyDayMARunner(generateBase())			// next step: should pass in buy/sell trigger structs
	fmt.Println(netStratReturn)

	// Graph Output
	http.HandleFunc("/", httpserver)
	http.ListenAndServe(":8081", nil)
	return
}

func main() {
	router := gin.Default()
	interval = 15 																	// 120 min interval (bot trigger freq)
	ticker = "TWTR"
	
	router.POST("/backtest/:ticker/:interval", postBacktest)
	router.Run("localhost:8050")
}