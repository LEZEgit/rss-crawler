package main

import "net/http"

// function signature you have to use if you want to 
// define an http handler in the way that the go std library expects

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	/*
		this just repsonds with a statusCode OK to tell that the server is alive and listening
	*/

	respondWithJSON(w, 200, struct{}{})
}