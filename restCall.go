package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
 * Example GET request
 */

func restGetCall(reqURL string) error {
	getReq, err := http.Get(reqURL)
	if err != nil {
		return err
	}
	if getReq.StatusCode != http.StatusOK {
		return fmt.Errorf("Unexpected status: %s", getReq.Status)
	}

	respBody, err := ioutil.ReadAll(getReq.Body)
	if err != nil {
		return err
	}

	var received response
	if err := json.Unmarshal(respBody, &received); err != nil {
		return err
	}

	fmt.Printf("response body: %s\n", respBody)
	fmt.Printf("message field value: %+v\n", received.Message)
	fmt.Printf("response status: %s\n", getReq.Status)

	return nil
}

/*
 * Example POST request
 */

// TODO: Shall be implemented
