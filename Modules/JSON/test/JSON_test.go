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

func TestEndPointGetWorld(t *testing.T) { // Testing the EndPointGet function from JSON package
	var worldObject JSON.WorldCases
	_, response := API.Availability(API.URLGenerator(Constants_and_Structs.WorldTotals, ""))
	println(response)
	body := API.ReadBody(response)
	println(string(body))
	worldObject = JSON.EndPointGetWorld(body)
	fmt.Println(worldObject)
	if worldObject.Results[0].Source.URL != "https://thevirustracker.com/" {
		t.Error("object empty")
	}
}

func TestEndPointGetCountry(t *testing.T) { // Testing the EndPointGet function from JSON package
	var countryObject JSON.CountryCases
	_, response := API.Availability(API.URLGenerator(Constants_and_Structs.CountryTotal, "TR"))
	println(response)
	body := API.ReadBody(response)
	println(string(body))
	countryObject = JSON.EndPointGetCountry(body)
	fmt.Println(countryObject)
	if countryObject.Countrydata[0].Info.Title != "Turkey" {
		t.Error("object empty")
	}
}

func TestCountryCalculator(t *testing.T) { // Testing the CountryCalculator function from JSON package
	var countryObject JSON.CountryCases
	_, response := API.Availability(API.URLGenerator(Constants_and_Structs.CountryTotal, "TR"))
	println(response)
	body := API.ReadBody(response)
	println(string(body))
	countryObject = JSON.EndPointGetCountry(body)
	fmt.Println(countryObject)
	if countryObject.Countrydata[0].Info.Title != "Turkey" {
		t.Error("Object empty")
	}
	var countryCalculatedVal = JSON.CountryCalculator(countryObject)
	fmt.Println(countryCalculatedVal)
	if countryCalculatedVal.TotalDeathPercentage != 0 {
		t.Errorf("Impossible Value")
	}
}

func TestWorldCalculator(t *testing.T) { // Testing the WorldCalculator function from JSON package
	var worldObject JSON.WorldCases
	_, response := API.Availability(API.URLGenerator(Constants_and_Structs.WorldTotals, ""))
	println(response)
	body := API.ReadBody(response)
	println(string(body))
	worldObject = JSON.EndPointGetWorld(body)
	fmt.Println(worldObject)
	if worldObject.Results[0].Source.URL != "https://thevirustracker.com/" {
		t.Error("object empty")
	}
	var worldCalculatedVal = JSON.WorldCalculator(worldObject)
	fmt.Println(worldCalculatedVal)
	if worldCalculatedVal.TotalDeathsPercentage != 0 {
		t.Errorf("Impossible Value")
	}
}
