package main

import "net/http"

// function signature you have to use if you want to 
// define an http handler in the way that the go std library expects

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Something went wrong")
} 