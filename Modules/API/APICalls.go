package API

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// This Package is made for API request and JSON parsing.

// *************** Global Variables *****************

var url = "https://pomber.github.io/covid19/timeseries.json"

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

func grabInformation(countryName string, dateGiven string) cases { // gives information about cases for requested
	// country and date.
	spaceClient := http.Client{
		Timeout: time.Second * 2, // Max of 2 secss
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "COVID-19_Analyser")

	res, getErr := spaceClient.Do(req)

	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)

	if readErr != nil {
		log.Fatal(readErr)
	}

	if body != nil {
		log.Fatal("Value of body is nil!")
	}

	situation := cases{}

	jsonError := json.Unmarshal(body, &situation)

	if jsonError != nil {
		log.Fatal(jsonError)
	}

	return situation
}

func APIAvaibility() bool {
	response, err := http.Get("https://pomber.github.io/covid19/timeseries.json")

	if err == nil && response != nil {
		return true
	} else {
		return false
	}
}
