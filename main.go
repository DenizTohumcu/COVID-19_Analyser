package main

import (
	"COVID-19_Analyser/Modules/API"
	"fmt"
)

func main() {
	API.HelloFunc()
	token, response := API.Avaibility("https://api.thevirustracker.com/free-api?global=stats")
	fmt.Println(token)
	fmt.Println(API.ReadBody(response))

}
