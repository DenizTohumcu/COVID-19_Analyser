package test

import (
	"COVID-19_Analyser/Modules/API"
	"COVID-19_Analyser/Modules/API/Constants_and_Structs"
	"COVID-19_Analyser/Modules/JSON"
	"testing"
)

// ***************** Test Struct Function ********************

var TestWorld JSON.WorldCases

// ***************** Global Variable ********************

var target Constants_and_Structs.URL

// ***************** Test functions ********************

func TestUnMarshalCases(t *testing.T) { // Testing the UnMarshal function from JSON package
	_, response := API.Availability(target.WorldTotals)
	body := API.ReadBody(response, "data")
	JSON.UnMarshalCases(body, &TestWorld)

}

func TestEndPointGet(t *testing.T) { // Testing the EndPointGet function from JSON package

}

func TestCountryCalculator(t *testing.T) { // Testing the CountryCalculator function from JSON package

}

func TestWorldCalculator(t *testing.T) { // Testing the WorldCalculator function from JSON package

}
