package API

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// This Package is made for API request and JSON parsing.

// *************** Global Variables *****************

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

type RegionCases struct {
}

type CountryCases struct {

	// End points JSON file of the API
	date      string `json:"date"`      // Date of information
	confirmed int    `json:"confirmed"` // Number of confirmed cases in a country
	deaths    int    `json:"deaths"`    // Number of deaths in a country per day
	recovered int    `json:"recovered"` // Number of recovered in a country per

	// Data calculated from the grabbed data
	closed             int //deaths + recovered // Number of closed cases in a country
	active             int //confirmed - closed // Number of active cases in a country
	deathPercentage    int //deaths/closed // Percentage of death in a country
	recoveryPercentage int //recovered/closed // Percentage of recovery in a country
	activePercentage   int //active/confirmed // Percentage of active cases in a country
	closedPercentage   int //closed/confirmed // Percentage of closed cases in a country

}

// *************** Function Declarations *******************
func HelloFunc() { // Hello initialisation function
	fmt.Println("Hello API!")
}

func Avaibility(url string) (bool, *http.Response) { // Returns true if API is
	response, err := http.Get(url)

	if err == nil && response != nil {
		return true, response
	} else {
		log.Fatal(err)
		return false, response
	}
}

func ReadBody(response *http.Response) string {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return string(data)
}
