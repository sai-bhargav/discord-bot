package util

import (
	"bytes"
	"log"
	"net/http"
)

// GetRequest is a generic method to make a GET requests
func GetRequest(url string) *http.Response {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("Failed to create GET request")
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")
	return makeRequest(req)
}

// PostRequest is a generic method to make a GET requests
func PostRequest(url string, body []byte, headers string) *http.Response {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		log.Printf("Failed to create POST request")
		panic(err)
	}

	if headers != "" {
		req.Header.Set("Content-Type", headers)
	} else {
		req.Header.Set("Content-Type", "application/json")
	}

	return makeRequest(req)
}

func makeRequest(req *http.Request) *http.Response {

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	return resp
}
