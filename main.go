package main

import (
	"COVID-19_Analyser/Modules/API"
	"fmt"
)

func main() {
	API.HelloFunc()
	fmt.Println(API.APIAvaibility())
}
