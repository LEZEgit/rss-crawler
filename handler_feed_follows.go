package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/LEZEgit/rss-crawler/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing request: %v", err))
		return
	}

	feed_follow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: params.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed_follow)) 
}

func (apiCfg *apiConfig) handlerUnfollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedIdStr := chi.URLParam(r, "feed_id")
	if feedIdStr == "" {
		respondWithError(w, 400, "Missing feed_id parameter")
		return
	}
	feedId, err := uuid.Parse(feedIdStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Invalid feed_id parameter: %v", err))
		return
	}

	err = apiCfg.DB.UnfollowFeed(r.Context(), database.UnfollowFeedParams{
		UserID: user.ID,
		FeedID: feedId,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't unfollow feed: %v", err))
		return
	}

	respondWithJSON(w, 200, map[string]string{"status": "unfollowed"})
}

func (apiCfg *apiConfig) handlerGetFeedsFollowingByUserId(w http.ResponseWriter, r *http.Request, user database.User) {
	// fetch the feeds that the user is following from the database
	feed_follows, err := apiCfg.DB.GetFeedsFollowingByUserID(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds following by user id: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedFollowsToFeedFollows(feed_follows))
}
