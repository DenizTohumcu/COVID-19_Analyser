package main

import (
	"fmt"

	API "./Modules/API"
)

func main() {
	API.HelloFunc()
	token, response := API.Avaibility("https://thevirustracker.com/timeline/map-data.json")
	fmt.Println(token)
	fmt.Println(API.ReadBody(response))
}
