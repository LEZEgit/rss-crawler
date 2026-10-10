package main

import (
	"net/http"
	"fmt"

	"github.com/LEZEgit/rss-crawler/internal/auth"
	"github.com/LEZEgit/rss-crawler/internal/database"
)

// this is the type of handler that requires authentication.
// It takes an additional parameter of type database.User, which is the authenticated user.
type authedHandler func(http.ResponseWriter, *http.Request, database.User)

// create a middleware function that takes an authedHandler and returns a standard http.HandlerFunc. 

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		// Get the API key from the request header
		apiKey, err := auth.GetAPIKeyFromRequest(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("API key error: %v", err))
			return
		}

		// find the user associated with the API key
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 404, fmt.Sprintf("User not found: %v", err))
			return
		}
		// Call the original handler with the authenticated user
		handler(w, r, user)
	}
}