package JSON

import (
	"encoding/json"
	"errors"
	"log"
)

// This module contains functions to parse and give demanded JSON files

// *************** Type Declarations *****************

type WorldCases struct {
	Results []struct {
		TotalCases             int `json:"total_cases"`
		TotalRecovered         int `json:"total_recovered"`
		TotalUnresolved        int `json:"total_unresolved"`
		TotalDeaths            int `json:"total_deaths"`
		TotalNewCasesToday     int `json:"total_new_cases_today"`
		TotalNewDeathsToday    int `json:"total_new_deaths_today"`
		TotalActiveCases       int `json:"total_active_cases"`
		TotalSeriousCases      int `json:"total_serious_cases"`
		TotalAffectedCountries int `json:"total_affected_countries"`
		Source                 struct {
			URL string `json:"url"`
		} `json:"source"`
	} `json:"results"`
	Stat string `json:"stat"`
}

type CountryCases struct {
	Countrydata []struct {
		Info struct {
			Ourid  int    `json:"ourid"`
			Title  string `json:"title"`
			Code   string `json:"code"`
			Source string `json:"source"`
		} `json:"info"`
		TotalCases          int `json:"total_cases"`
		TotalRecovered      int `json:"total_recovered"`
		TotalUnresolved     int `json:"total_unresolved"`
		TotalDeaths         int `json:"total_deaths"`
		TotalNewCasesToday  int `json:"total_new_cases_today"`
		TotalNewDeathsToday int `json:"total_new_deaths_today"`
		TotalActiveCases    int `json:"total_active_cases"`
		TotalSeriousCases   int `json:"total_serious_cases"`
		TotalDangerRank     int `json:"total_danger_rank"`
	} `json:"countrydata"`
	Stat string `json:"stat"`
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

// ######## Endpoint getter function ########

func EndPointGet(Name string, body []byte) (interface{}, error) { // Gets the end point values and

	if Name == "World" || Name == "WorldCases" {
		object := WorldCases{}
		json.Unmarshal(body, &object.totalCases)
		json.Unmarshal(body, &object.totalRecovered)
		json.Unmarshal(body, &object.totalUnresolved)
		json.Unmarshal(body, &object.totalDeaths)
		json.Unmarshal(body, &object.totalNewCasesToday)
		json.Unmarshal(body, &object.totalNewDeathsToday)
		json.Unmarshal(body, &object.totalActiveCases)
		json.Unmarshal(body, &object.totalSeriousCases)
		json.Unmarshal(body, &object.totalAffectedCountries)
		return object, nil
	} else if Name == "Country" || Name == "CountryCases" {
		object := CountryCases{}
		json.Unmarshal(body, &object.countryCode)
		json.Unmarshal(body, &object.date)
		json.Unmarshal(body, &object.cases)
		json.Unmarshal(body, &object.deaths)
		json.Unmarshal(body, &object.recovered)

		return object, nil
	} else {
		log.Fatal("Constructor error: given name is not found")
		return "Given name is not found : JSON.go -> line 103 ", errors.New("Given name is not found : JSON.go -> line 103 ")
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
