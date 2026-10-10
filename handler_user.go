package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/LEZEgit/rss-crawler/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing request: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name: params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUsertoUser(user)) 
}

func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		APIKey string `json:"api_key"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing request: %v", err))
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), params.APIKey)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user by API key: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUsertoUser(user)) 
}