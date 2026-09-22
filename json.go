package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	/*
		Format the msg into a consistent JSON object
	*/
	if statusCode > 499 { // Server side error
		log.Println("Responding with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondWithJSON(w, statusCode, errResponse {
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload any) {
	/*
		This function takes in the payload, and 
		adds it to the Response as a JSON object

		Marhsal does Go data -> JSON object
	*/



	// marshal the payload (Golang data) into a JSON object/string
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to Marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}