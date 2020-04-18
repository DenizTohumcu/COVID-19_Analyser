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

type cases struct {

	// Data grabbed from JSON file of the API
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

func ReadBody(response *http.Response) []byte {
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return data
}
