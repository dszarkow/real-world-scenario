package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
)

/*var (
	lightbulbs = make(map[string]bool)
)*/

/*
After running the main function, we're using the HTTP package to do a GET HTTP call
to our service on the lightbulbs endpoint. The result of the http.Get function is an
HTTP response and an error. The first thing we do is check for the error. Then,
we check if the response status code is what we expected. Finally, we check
if the response body is what we expected.
*/
func TestLightbulbs(t *testing.T) {
	go main()
	response, err := http.Get("http://localhost:8080/lightbulbs")
	if err != nil {
		t.Errorf("expected no errors, but got %v", err)
	}

	if response.StatusCode != http.StatusOK {
		t.Errorf("expected 200 statuscode, but got %v", response.StatusCode)
	}

	//responseBody := make(map[string]interface{})
	//json.NewDecoder(response.Body).Decode(&responseBody)
	json.NewDecoder(response.Body).Decode(&lightbulbs)
	fmt.Println(lightbulbs)
	response.Body.Close()

	if lightbulbs == nil {
		t.Errorf(`expected lightbulbs map", but got %v`, lightbulbs)
	}

	os.Interrupt.Signal()
}
