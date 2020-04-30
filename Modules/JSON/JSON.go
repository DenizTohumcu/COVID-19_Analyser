package JSON

import (
	"encoding/json"
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
	TotalClosed int // deaths + recovered // Number of closed cases in a country
	//active                    int // confirmed - closed // Number of active cases in a country
	TotalDeathPercentage        int // deaths/closed      // Percentage of death in a country
	TotalRecoveryPercentage     int // recovered/closed   // Percentage of recovery in a country
	TotalActivePercentage       int // active/confirmed   // Percentage of active cases in a country
	TotalClosedPercentage       int // closed/confirmed   // Percentage of closed cases in a country
	TotalDeficit                int // totalNewCasesToday - totalNewDeathsToday // Value that show the change in trends
	TotalSeriousCasesPercentage int // totalSeriousCases/totalActiveCases       // Percentage of Serious cases "
	TotalActiveCasesPercentage  int // totalActiveCases/totalCases              // Percentage of active cases "
}

type WorldCalculatedValues struct {

	// Data calculated from the grabbed data with WorldCases object
	TotalRecoveredPercentage    int // totalRecovered/totalCases                // Percentage of recovery in the World
	TotalUnresolvedPercentage   int // totalUnresolved/totalCases               // Percentage of death "
	TotalClosed                 int // totalCases - activeCases                 // Number of closed cases "
	TotalDeathsPercentage       int // totalDeaths/totalClosed                  // Percentage of deaths over closed "
	TotalActivePercentage       int // active/confirmed                         // Percentage of active cases "
	TotalClosedPercentage       int // closed/confirmed                         // Percentage of closed cases "
	TotalDeficit                int // totalNewCasesToday - totalNewDeathsToday // Value that show the change in trends
	TotalSeriousCasesPercentage int // totalSeriousCases/totalActiveCases       // Percentage of Serious cases "

}

// *************** Function Declarations *******************

// ######## Endpoint getter function ########

func EndPointGetWorld(body []byte) WorldCases { // Gets the end point values and

	object := WorldCases{}
	err := json.Unmarshal(body, &object)
	if err != nil {
		log.Fatal(err)
	}

	return object
}

func EndPointGetCountry(body []byte) CountryCases {

	object := CountryCases{}
	err := json.Unmarshal(body, &object)
	if err != nil {
		log.Fatal(err)
	}

	return object

}

// ######## Information handling function ########

func CountryCalculator(country CountryCases) CountryCalculatedValues { // Calculates the given values by the API for
	// each country
	totalClosed := country.Countrydata[0].TotalRecovered + country.Countrydata[0].TotalCases
	totalDeathPercentage := (country.Countrydata[0].TotalDeaths / totalClosed) * 100
	totalRecoveryPercentage := (country.Countrydata[0].TotalRecovered / totalClosed) * 100
	totalActivePercentage := (country.Countrydata[0].TotalActiveCases / country.Countrydata[0].TotalCases) * 100
	totalClosedPercentage := (totalClosed / country.Countrydata[0].TotalCases) * 100
	totalDeficit := country.Countrydata[0].TotalNewCasesToday - country.Countrydata[0].TotalNewDeathsToday
	totalSeriousCasesPercentage := country.Countrydata[0].TotalSeriousCases / country.Countrydata[0].TotalActiveCases
	totalActiveCasesPercentage := country.Countrydata[0].TotalActiveCases / country.Countrydata[0].TotalCases

	return CountryCalculatedValues{totalClosed, totalDeathPercentage,
		totalRecoveryPercentage, totalActivePercentage,
		totalClosedPercentage, totalDeficit,
		totalSeriousCasesPercentage, totalActiveCasesPercentage}
}

func WorldCalculator(world WorldCases) WorldCalculatedValues { //Calculates the given values by the API for the world

	totalRecoveredPercentage := (world.Results[0].TotalRecovered / world.Results[0].TotalCases) * 100
	totalUnresolvedPercentage := (world.Results[0].TotalUnresolved / world.Results[0].TotalCases) * 100
	totalClosed := world.Results[0].TotalCases - world.Results[0].TotalUnresolved
	totalDeathsPercentage := (world.Results[0].TotalDeaths / totalClosed) * 100
	totalActivePercentage := (world.Results[0].TotalActiveCases / world.Results[0].TotalCases) * 100
	totalClosedPercentage := (totalClosed / world.Results[0].TotalCases) * 100
	totalDeficit := world.Results[0].TotalNewCasesToday - world.Results[0].TotalNewDeathsToday
	totalSeriousCasesPercentage := (world.Results[0].TotalSeriousCases / world.Results[0].TotalActiveCases) * 100

	return WorldCalculatedValues{totalRecoveredPercentage,
		totalUnresolvedPercentage, totalClosed,
		totalDeathsPercentage, totalActivePercentage,
		totalClosedPercentage, totalDeficit,
		totalSeriousCasesPercentage}
}
