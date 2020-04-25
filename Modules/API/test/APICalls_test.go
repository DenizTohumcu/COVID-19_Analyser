package main

import (
	"COVID-19_Analyser/Modules/API"
	"testing"
)

// ***************** Test functions ********************

func TestAvailability(t *testing.T) { // Testing the Availability function from API package
	token, response := API.Availability("https://thevirustracker.com/timeline/map-data.json")
	if token != true && response.Status != "Accepted" { // Waiting for True and Accepted
		t.Errorf("Sum was incorrect, got: %t, want: %t.", token, true)
	}
}

func TestReadbody(t *testing.T) { // Testing for body reading function
	token, response := API.Availability("https://thevirustracker.com/timeline/map-data.json")

	if API.ReadBody(response, "string") != "" && token != true { // Waiting for non-empty string and true
		t.Errorf("The body is empty.")

	}
}
