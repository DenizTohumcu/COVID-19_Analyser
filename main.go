package main

import (
	"COVID-19_Analyser/Modules/API"
	"fmt"
)

func main() {
	API.HelloFunc()
	token, response := API.Avaibility("https://thevirustracker.com/timeline/map-data.json")
	fmt.Println(token)
	fmt.Println(API.ReadBody(response))
}
