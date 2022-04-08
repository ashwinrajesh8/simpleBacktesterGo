package main


import (
	"net/http"
	"io/ioutil"
    "log"
	"fmt"
	"encoding/json"
	"strconv"
)


func generateBase() *AggregateJSON {
	resp, err := http.Get("https://api.polygon.io/v2/aggs/ticker/"+ticker+"/range/"+strconv.Itoa(interval)+"/minute/"+TimeIncrementor+"/"+Timestamp+"?adjusted=true&sort=asc&apiKey="+apiKey)
	if err != nil {
	   log.Fatalln(err)
	}
	 //We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
	   log.Fatalln(err)
	}
 //Convert the body to type string
	sb := string(body)
	fmt.Println(sb)
 
	//json stuff
	aggregateJSON := &AggregateJSON{}
	errTwo := json.Unmarshal([]byte(sb), aggregateJSON)

	if errTwo == nil {
		fmt.Println(errTwo)
		fmt.Printf("\nFAIL\n")
	}
	fmt.Printf("\n\n ticker: %+v\n", aggregateJSON.Ticker)
	return aggregateJSON
}