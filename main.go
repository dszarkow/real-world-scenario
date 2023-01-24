// Ref: https://x-team.com/blog/go-crash-course-2/
//
// Modified by Don Szarkowicz, January 2023

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Create a simple map to hold the state of the bulbs in our house.
var (
	lightbulbs = make(map[string]bool)
)

func main() {
	//Assign a few standard bulbs to the simple map
	lightbulbs["livingroom"] = false
	lightbulbs["kitchen"] = false

	// Create an HTTP handler by binding a path named /hello to a function.
	// The HTTP handler functions will always receive two parameters:
	// 1) http.ResponseWriter to write a response back to whoever requested it.
	// 2) *http.Request to understand the type of request we're dealing with.
	//
	// Usage: http://localhost:8080/hello
	http.HandleFunc("/hello", func(responseWriter http.ResponseWriter, request *http.Request) {
		// Set the HTML header content type
		responseWriter.Header().Set("Content-Type", "application-json")
		// Set the HTML header status to StatusOK
		responseWriter.WriteHeader(http.StatusOK)
		// Handwrite a JSON string and convert it to []byte so it can be used in the Write function
		responseWriter.Write([]byte(`{"message":"Hello World!"}`))
	})

	// Create an HTTP handler by binding a path named /healthcheck to a function.
	// The HTTP handler functions will always receive two parameters:
	// 1) http.ResponseWriter to write a response back to whoever requested it.
	// 2) *http.Request to understand the type of request we're dealing with.
	//
	// Usage: http://localhost:8080/healthcheck
	http.HandleFunc("/healthcheck", func(responseWriter http.ResponseWriter, request *http.Request) {
		// Set the HTML header content type
		responseWriter.Header().Set("Content-Type", "application-json")
		// Set the HTML header status to StatusOK
		responseWriter.WriteHeader(http.StatusOK)
		// Handwrite a JSON string and convert it to []byte so it can be used in the Write function
		responseWriter.Write([]byte(`{"message":"service is up and running"}`))
	})

	// Create an HTTP handler by binding a path named /lightbulbs to a function.
	// The HTTP handler functions will always receive two parameters:
	// 1) http.ResponseWriter to write a response back to whoever requested it.
	// 2) *http.Request to understand the type of request we're dealing with.
	//
	// This endpoint will output the current state of our bulb map as Json data.
	//
	// Usage: http://localhost:8080/lightbulbs
	http.HandleFunc("/lightbulbs", func(responseWriter http.ResponseWriter, request *http.Request) {
		// Set the HTML header content type
		responseWriter.Header().Set("Content-Type", "application-json")
		// Set the HTML header status to StatusOK
		responseWriter.WriteHeader(http.StatusOK)
		// Encode the lightbulbs map as a JSON string and pass it to the responseWriter
		json.NewEncoder(responseWriter).Encode(lightbulbs)
	})

	// Create an HTTP handler by binding a path named /lightbulbs/switch to a function.
	// The HTTP handler functions will always receive two parameters:
	// 1) http.ResponseWriter to write a response back to whoever requested it.
	// 2) *http.Request to understand the type of request we're dealing with.
	//
	// This endpoint should toggle the boolean value of a named lightbulb,
	// then output the current state of our bulb map as Json data.
	//
	// Usage: http://localhost:8080/lightbulbs/switch?name=kitchen or name=livingroom
	http.HandleFunc("/lightbulbs/switch", func(responseWriter http.ResponseWriter, request *http.Request) {
		// Set the HTML header content type
		responseWriter.Header().Set("Content-Type", "application-json")

		// Get the value of the name parameter
		name := request.URL.Query().Get("name")

		// If no name was entered...
		if name == "" {
			// Set the header status to StatusBadRequest
			responseWriter.WriteHeader(http.StatusBadRequest)
			// Handwrite a JSON string and convert it to []byte so it can be used in the Write function
			responseWriter.Write([]byte(`{"message":"a light bulb name should be provided as the value of a
		'name' querystring parameter"}`))
			// Return from this function
			return
		}

		// If the entered name is not a lightbulbs map key...
		if _, keyExists := lightbulbs[name]; !keyExists {
			// Set the header status to StatusNotFound
			responseWriter.WriteHeader(http.StatusNotFound)
			// Handwrite a JSON string and convert it to []byte so it can be used in the Write function
			responseWriter.Write([]byte(`{"message":"a light bulb with the provided name doesn't exist"}`))
			// Return from this function
			return
		}

		// Here the name is a valid key so toggle the value
		lightbulbs[name] = !lightbulbs[name]

		// Set the HTML header status to StatusOK
		responseWriter.WriteHeader(http.StatusOK)

		// Encode the lightbulbs map as a JSON string and pass it to the responseWriter
		json.NewEncoder(responseWriter).Encode(lightbulbs)
	})

	// Create an HTTP handler by binding a path named /lightbulbs/switch to a function.
	// The HTTP handler functions will always receive two parameters:
	// 1) http.ResponseWriter to write a response back to whoever requested it.
	// 2) *http.Request to understand the type of request we're dealing with.
	//
	// This endpoint should create a new lightbulb map entry
	// then output the current state of our bulb map as Json data.
	//
	// Usage: http://localhost:8080/lightbulbs/create?name=bedroom
	http.HandleFunc("/lightbulbs/create", func(responseWriter http.ResponseWriter, request *http.Request) {
		// Set the HTML header content type
		responseWriter.Header().Set("Content-Type", "application-json")

		// Get the value of the name parameter
		name := request.URL.Query().Get("name")

		// If no name was entered...
		if name == "" {
			// Set the header status to StatusBadRequest
			responseWriter.WriteHeader(http.StatusBadRequest)

			// Handwrite a JSON string and convert it to []byte so it can be used in the Write function
			responseWriter.Write([]byte(`{"message":"a light bulb name should be provided as the value of a
		'name' querystring parameter"}`))

			// Return from this function
			return
		}

		// If the entered name is already a lightbulbs map key...
		if _, keyExists := lightbulbs[name]; keyExists {
			// Set the header status to StatusBadRequest
			responseWriter.WriteHeader(http.StatusBadRequest)

			// Handwrite a JSON string and convert it to []byte so it can be used in the Write function
			responseWriter.Write([]byte(`{"message":"a lightbulb with the provided name already exists"}`))

			// Return from this function
			return
		}

		// Here the name is a not currently a valid key, so create a new map entry with value false
		lightbulbs[name] = false

		// Set the header status to StatusOK
		responseWriter.WriteHeader(http.StatusOK)

		// Encode the lightbulbs map as a JSON string and pass it to the responseWriter
		json.NewEncoder(responseWriter).Encode(lightbulbs)
	})

	// Usage: http://localhost:8080/lightbulbs/delete?name=bedroom
	http.HandleFunc("/lightbulbs/delete", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Header().Set("Content-Type", "application-json")
		name := request.URL.Query().Get("name")
		if name == "" {
			responseWriter.WriteHeader(http.StatusBadRequest)
			responseWriter.Write([]byte(`{"message":"a light bulb name should be provided as the value of a
		'name' querystring parameter"}`))
			return
		}
		if _, keyExists := lightbulbs[name]; !keyExists {
			responseWriter.WriteHeader(http.StatusNotFound)
			responseWriter.Write([]byte(`{"message":"a lightbulb with the provided name doesn't exist"}`))
			return
		}
		delete(lightbulbs, name)
		responseWriter.WriteHeader(http.StatusOK)
		json.NewEncoder(responseWriter).Encode(lightbulbs)
	})

	// Print a message to the terminal so we know what's happening
	fmt.Println("http server listening on localhost:8080")

	// Start this server on port 8080
	// http.HandleFunc(), used in the handler functions above,
	// binds that handler to the default HTTP handler.
	// When http.ListenAndServe doesn't receive any handlers (i.e. nil is passed),
	// it uses that default HTTP handler.
	http.ListenAndServe(":8080", nil)
}
