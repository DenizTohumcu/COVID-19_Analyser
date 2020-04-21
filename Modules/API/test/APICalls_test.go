package main

import (
	".."
	"testing"
)

// ***************** Test functions ********************

func TestAvaibility(t *testing.T) {
	token, response := API.Avaibility("https://thevirustracker.com/timeline/map-data.json")
	if token != true && response.Status != "Accepted" {
		t.Errorf("Sum was incorrect, got: %t, want: %t.", token, true)
	}
}

func TestReadbody(t *testing.T) {
	token, response := API.Avaibility("https://thevirustracker.com/timeline/map-data.json")

	if API.ReadBody(response) != "" && token != true {
		t.Errorf("The body is empty.")

	}
}
