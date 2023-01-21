package main

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"
)

func TestHealthCheck(t *testing.T) {
	go main()
	response, err := http.Get("http://localhost:8080/healthcheck")
	if err != nil {
		t.Errorf("expected no errors, but got %v", err)
	}
	if response.StatusCode != http.StatusOK {
		t.Errorf("expected 200 statuscode, but got %v", response.StatusCode)
	}
	responseBody := make(map[string]interface{})
	json.NewDecoder(response.Body).Decode(&responseBody)
	response.Body.Close()
	if responseBody["message"] != "service is up and running" {
		t.Errorf(`expected message to be "service is up and running", but got %v`, responseBody["message"])
	}
	os.Interrupt.Signal()
}
