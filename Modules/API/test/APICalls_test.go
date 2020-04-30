package main

import (
	"COVID-19_Analyser/Modules/API"
	"COVID-19_Analyser/Modules/API/Constants_and_Structs"
	"fmt"
	"testing"
)

// ***************** Test functions ********************

func TestAvailability(t *testing.T) { // Testing the Availability function from API package
	token, response := API.Availability("https://thevirustracker.com/timeline/map-data.json")
	if token != true && response.Status != "Accepted" { // Waiting for True and Accepted
		t.Errorf("Sum was incorrect, got: %t, want: %t.", token, true)
	}
	fmt.Println(response)
}

func TestReadbody(t *testing.T) { // Testing for body reading function
	token, response := API.Availability("https://thevirustracker.com/timeline/map-data.json")
	data := API.ReadBody(response)
	fmt.Println(string(data))
	if data != nil && token != true { // Waiting for non-empty string and true
		t.Errorf("The body is empty.")
	}
}

func TestURLGenerator(t *testing.T) {

	ReqAddress := API.URLGenerator(Constants_and_Structs.CountryTotal, "TR")
	fmt.Println(ReqAddress)
	if ReqAddress != "https://api.thevirustracker.com/free-api?countryTotal=TR" {
		t.Errorf("Expected URL not generated.")
	}
}
