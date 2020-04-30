package test

import (
	"COVID-19_Analyser/Modules/API"
	"COVID-19_Analyser/Modules/API/Constants_and_Structs"
	"COVID-19_Analyser/Modules/JSON"
	"fmt"
	"testing"
)

// ***************** Test Struct Function ********************

// ***************** Global Variable ********************

// ***************** Test functions ********************

func TestEndPointGet(t *testing.T) { // Testing the EndPointGet function from JSON package
	var countryObject JSON.CountryCases
	_, response := API.Availability(API.URLGenerator(Constants_and_Structs.CountryTotal, "TR"))
	println(response)
	body := API.ReadBody(response)
	println(string(body))
	countryObject = JSON.EndPointGet("Country", body)
	fmt.Println(countryObject)
	if countryObject != nil {
		t.Error("object empty")
	}
}

func TestCountryCalculator(t *testing.T) { // Testing the CountryCalculator function from JSON package

}

func TestWorldCalculator(t *testing.T) { // Testing the WorldCalculator function from JSON package

}
