package test

import (
	"COVID-19_Analyser/Modules/API"
	"COVID-19_Analyser/Modules/API/Constants_and_Structs"
	"COVID-19_Analyser/Modules/JSON"
	"fmt"
	"testing"
)

// ***************** Test Struct Function ********************

var TestWorld JSON.WorldCases

// ***************** Global Variable ********************

var target Constants_and_Structs.URL

// ***************** Test functions ********************

func TestEndPointGet(t *testing.T) { // Testing the EndPointGet function from JSON package
	_, response := API.Availability(target.WorldTotals)
	println(response)
	body := API.ReadBody(response)
	println(string(body))
	_, err := JSON.EndPointGet("World", body)
	fmt.Println(TestWorld)
	if err != nil {
		t.Error(err)
	}
}

func TestCountryCalculator(t *testing.T) { // Testing the CountryCalculator function from JSON package

}

func TestWorldCalculator(t *testing.T) { // Testing the WorldCalculator function from JSON package

}
