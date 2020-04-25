package JSON

import (
	"encoding/json"
	"log"
)

// This module contains functions to parse and give demanded JSON files

// *************** Type Declarations *****************

type WorldCases struct {
	// End points JSON file of the API
	totalCases             int `json:"total_cases"`              // Total cases in the world
	totalRecovered         int `json:"total_recovered"`          // Total recovered cases in the world
	totalUnresolved        int `json:"total_unresolved"`         // Total unresolved cases in the world
	totalDeaths            int `json:"total_deaths"`             // Total death cases in the world
	totalNewCasesToday     int `json:"total_new_cases_today"`    // Total new cases today in the world
	totalNewDeathsToday    int `json:"total_new_deaths_today"`   // Total new deaths in the world
	totalActiveCases       int `json:"total_active_cases"`       // Total active cases in the world
	totalSeriousCases      int `json:"total_serious_cases"`      // Total serious cases in the world
	totalAffectedCountries int `json:"total_affected_countries"` // Total affected cases in the world

}

type CountryCases struct {

	// End points JSON file of the API
	countryCode string `json:"countrycode"` // Date of information
	date        string `json:"date"`        // Date of information
	cases       int    `json:"cases"`       // Number of confirmed cases in a country
	deaths      int    `json:"deaths"`      // Number of deaths in a country per day
	recovered   int    `json:"recovered"`   // Number of recovered in a country per

}

type CountryCalculatedValues struct {

	// Data calculated from the grabbed data with CountryCases object
	closed             int // deaths + recovered // Number of closed cases in a country
	active             int // confirmed - closed // Number of active cases in a country
	deathPercentage    int // deaths/closed      // Percentage of death in a country
	recoveryPercentage int // recovered/closed   // Percentage of recovery in a country
	activePercentage   int // active/confirmed   // Percentage of active cases in a country
	closedPercentage   int // closed/confirmed   // Percentage of closed cases in a country

}

type WorldCalculatedValues struct {

	// Data calculated from the grabbed data with WorldCases object
	totalRecoveredPercentage    int // totalRecovered/totalCases                // Percentage of recovery in the World
	totalUnresolvedPercentage   int // totalUnresolved/totalCases               // Percentage of death "
	totalClosed                 int // totalCases - activeCases                 // Number of closed cases "
	totalDeathsPercentage       int // totalDeaths/totalClosed                  // Percentage of deaths over closed "
	totalDeficit                int // totalNewCasesToday - totalNewDeathsToday // Value that show the change in trends
	totalSeriousCasesPercentage int // totalSeriousCases/totalActiveCases       // Percentage of Serious cases "
	totalActiveCasesPercentage  int // totalActiveCases/totalCases              // Percentage of active cases "

}

// *************** Function Declarations *******************

// ######## JSON related functions #########

func UnMarshal(body []byte, key interface{}) interface{} { //json unmarshal with err handling

	err := json.Unmarshal(body, key)
	if err != nil {
		log.Fatal("Decoding error: ", err)
	}

	return key
}

// ######## Endpoint getter function ########

func EndpointGet(Name string, body []byte) interface{} {

	if Name == "World" || Name == "WorldCases" {
		object := WorldCases{}
		UnMarshal(body, object.totalCases)
		UnMarshal(body, object.totalRecovered)
		UnMarshal(body, object.totalUnresolved)
		UnMarshal(body, object.totalDeaths)
		UnMarshal(body, object.totalNewCasesToday)
		UnMarshal(body, object.totalNewDeathsToday)
		UnMarshal(body, object.totalActiveCases)
		UnMarshal(body, object.totalSeriousCases)
		UnMarshal(body, object.totalAffectedCountries)
		return object
	} else if Name == "Country" || Name == "CountryCases" {
		object := CountryCases{}
		UnMarshal(body, object.countryCode)
		UnMarshal(body, object.date)
		UnMarshal(body, object.cases)
		UnMarshal(body, object.deaths)
		UnMarshal(body, object.recovered)

		return object
	} else {
		log.Fatal("Constructor error: given name is not found")
		return "Given name is not found : JSON.go -> line 103 "
	}

}

// ######## Information handling function ########

func CountryCalculator(country CountryCases) CountryCalculatedValues { // Calculates the given values by the API for
	// each country
	closed := country.deaths + country.recovered
	active := country.cases - closed
	deathPercentage := (country.deaths / closed) * 100
	recoveryPercentage := (country.recovered / closed) * 100
	activePercentage := (active / country.cases) * 100
	closedPercentage := (closed / country.cases) * 100

	return CountryCalculatedValues{closed, active, deathPercentage,
		recoveryPercentage, activePercentage, closedPercentage}
}

func WorldCalculator(world WorldCases) WorldCalculatedValues { //Calculates the given values by the API for the world

	totalRecoveredPercentage := (world.totalRecovered / world.totalCases) * 100
	totalUnresolvedPercentage := (world.totalUnresolved / world.totalCases) * 100
	totalClosed := world.totalCases - world.totalUnresolved
	totalDeathsPercentage := (world.totalDeaths / totalClosed) * 100
	totalDeficit := world.totalNewCasesToday - world.totalNewDeathsToday
	totalSeriousCasesPercentage := world.totalSeriousCases / world.totalActiveCases
	totalActiveCasesPercentage := world.totalActiveCases / world.totalCases

	return WorldCalculatedValues{totalRecoveredPercentage,
		totalUnresolvedPercentage, totalClosed,
		totalDeathsPercentage, totalDeficit,
		totalSeriousCasesPercentage, totalActiveCasesPercentage}
}
