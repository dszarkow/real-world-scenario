package main

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
)

/*var (
	lightbulbs = make(map[string]bool)
)*/

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
	response.Body.Close()

	if lightbulbs == nil {
		t.Errorf(`expected lightbulbs map", but got %v`, lightbulbs)
	}

	os.Interrupt.Signal()
}
