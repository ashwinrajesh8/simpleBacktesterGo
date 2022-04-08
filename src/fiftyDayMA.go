package main

import (
	"net/http"
	"io/ioutil"
    "log"
	"fmt"
	"encoding/json"
	"time"
)

func fiftyDayMovingAverage(ticky string, apiKey string) float64 {
	respTwo, errFour := http.Get("https://api.polygon.io/v2/aggs/ticker/"+ticky+"/range/1/day/"+timey.AddDate(0,0,-72).Format("2006-01-02")+"/"+timey.AddDate(0,0,-1).Format("2006-01-02")+"?adjusted=true&sort=asc&limit=120&apiKey="+apiKey)		// was -75, replace time.Now() with timey
	fmt.Println(time.Now().AddDate(0,0,-1).Format("2006-01-02"))
	fmt.Println(time.Now().AddDate(0,0,-72).Format("2006-01-02"))  
  if errFour != nil {
	 log.Fatalln(errFour)
  }
   //We Read the response body on the line below.
  bodyTwo, errFour := ioutil.ReadAll(respTwo.Body)
  if errFour != nil {
	 log.Fatalln(errFour)
  }
//Convert the body to type string
  sbTwo := string(bodyTwo)
  fmt.Println(sbTwo)
  aggregateJSONtwo := &AggregateJSON{}
  errTree := json.Unmarshal([]byte(sbTwo), aggregateJSONtwo)
//    var getPoly polygonRequest
//    errTwo := json.Unmarshal([]byte(sb), &getPoly)
  if errTree == nil {
	  fmt.Println(errTree)
	  fmt.Printf("\nFAIL\n")
  }
  closeSum := 0.0000
  fmt.Printf("\n\n ticker: %+v\n", aggregateJSONtwo.Ticker)
  dayCounter := 0
  for i := 0; i < len(aggregateJSONtwo.Results); i++ {
	  fmt.Printf("Summer day %d: %+v\n", i+1, aggregateJSONtwo.Results[i].C)
	  closeSum = closeSum + aggregateJSONtwo.Results[i].C
	  dayCounter++
  }
  fiftyDayMovingAvg := closeSum/float64(dayCounter) // just in case not exactly 50 days	
  fmt.Printf("Fifty Day Moving Average: %f", fiftyDayMovingAvg)
  return fiftyDayMovingAvg
}

func fiftyDayMARunner(aggregateJSON *AggregateJSON) float64 {	// takes in json, buy/sell, crosses/is above/below
	startingAmount := 5000.000000
	purchasePower := 5000.000000
	portfolioVal := 0.000000
	holdingNum := 0
	PnL := 0.000000
	currQuote := 0.000000
	startVal := aggregateJSON.Results[0].C
	trackDay := 0
	fiftyDayMovingAvg := fiftyDayMovingAverage(ticker, apiKey)
	for i := 0; i < len(aggregateJSON.Results); i++ {
		day := time.UnixMilli(aggregateJSON.Results[i].T).Day()
		if(trackDay != int(day)){
		 timey = time.UnixMilli(aggregateJSON.Results[i].T)
		 fiftyDayMovingAvg = fiftyDayMovingAverage(ticker, apiKey)
		 trackDay = int(day)
		}
		// INSERT Strat evals here
		if(aggregateJSON.Results[i].C>fiftyDayMovingAvg && purchasePower>aggregateJSON.Results[i].C){
			 holdingNum++
			 purchasePower = purchasePower - aggregateJSON.Results[i].C
		}
		 portfolioVal += aggregateJSON.Results[i].C
		 currQuote = aggregateJSON.Results[i].C
			fmt.Printf("Quotes: %+v  %f %f\n", aggregateJSON.Results[i].C, purchasePower, float64(holdingNum)*currQuote+purchasePower-startingAmount)
		performance = append(performance, results{aggregateJSON.Results[i].C-startVal, float64(holdingNum)*currQuote+purchasePower-startingAmount})
		dataPoints += 1
	}
	presentVal := float64(holdingNum) * currQuote
	PnL = presentVal + purchasePower - startingAmount
	fmt.Println(PnL)

	return PnL
}