package main

import (
	"encoding/json"
	"io/ioutil"
)

func readParams() (url string, err error) {
	dataFile, err := ioutil.ReadFile("parameters.json")
	if err != nil {
		return "", err
	}

	var fParams fileParams
	if err := json.Unmarshal(dataFile, &fParams); err != nil {
		return "", err
	}

	var paramURL string = "http://" + fParams.IP + ":8443"

	return paramURL, nil
}
