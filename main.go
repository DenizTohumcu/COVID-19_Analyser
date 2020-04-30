package main

import (
	"COVID-19_Analyser/Modules/API"
	"encoding/json"
	"fmt"
)

func main() {
	API.HelloFunc()
	token, response := API.Availability("https://api.thevirustracker.com/free-api?countryTotal=TR")
	fmt.Println(token)
	//fmt.Println(API.ReadBody(response))
	CountryType := API.ReadBody(response)
	//TotalCases := json.Unmarshal(CountryType,JSON.)
	//fmt.Println(TotalCases)
}
