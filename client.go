package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/mitchellh/mapstructure"
)

// APIClient for openov api
type APIClient struct {
	Host string
}

// NewAPIClient creates, caches and returns a new api client
func NewAPIClient() *APIClient {
	client := &APIClient{Host: "http://v0.ovapi.nl/"}
	return client
}

// GetDepartures of a stop area
func (a *APIClient) GetDepartures(stopAreaCode string) (*Departure, error) {
	var reqURL = a.Host + "stopareacode/" + stopAreaCode + "/departures"
	jsonBlob, err := a.get(reqURL)

	var stopArea map[string]interface{}
	err = json.Unmarshal(jsonBlob, &stopArea)

	if err != nil {
		return nil, err
	}

	departures := NewDeparture()
	var timingCode string

	// loop over stop areas
	for _, stopAreaValue := range stopArea {
		// loop over timingpoints
		for timingPointKey, timingPointValue := range stopAreaValue.(map[string]interface{}) {
			if timingCode == "" {
				timingCode = timingPointKey
			}
			// from the timing point get the passes
			for timingPointGroupKey, timingPointGroupValue := range timingPointValue.(map[string]interface{}) {
				if timingPointGroupKey != "Passes" {
					continue
				}
				for _, pass := range timingPointGroupValue.(map[string]interface{}) {
					var timingPointPass TimingPointPass
					mapstructure.Decode(pass, &timingPointPass)
					if timingPointPass.TimingPointCode == timingCode {
						departures.AddTimingPointPass(timingPointPass, true)
					} else {
						departures.AddTimingPointPass(timingPointPass, false)
					}
				}
			}
		}
	}

	return departures, nil
}

func (a *APIClient) get(path string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", path, nil)
	req.Header.Add("Content-Encoding", "gzip")
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// look at go wow api for inspiration:
// https://github.com/capoferro/wow/blob/master/api_client.go
