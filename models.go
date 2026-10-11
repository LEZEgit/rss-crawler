package main

import (
	"time"
	"uuid"

	"github.com/LEZEgit/rss-crawler/internal/database"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	ApiKey	string `json:"api_key"`
}

func databaseUsertoUser(dbUser database.User) User {
	return User {
		ID: uuid.UUID(dbUser.ID),
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name: dbUser.Name,
		ApiKey: dbUser.ApiKey,
	}
}


type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string `json:"name"`
	Url       string `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func databaseFeedtoFeed(dbFeed database.Feed) Feed {
	return Feed {
		ID: uuid.UUID(dbFeed.ID),
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name: dbFeed.Name,
		Url: dbFeed.Url,
		UserID: uuid.UUID(dbFeed.UserID),
	}
}



func databaseFeedsToFeeds(dbFeeds []database.Feed) []Feed {
	feeds := make([]Feed, len(dbFeeds))
	for i, dbFeed := range dbFeeds {
		feeds[i] = databaseFeedtoFeed(dbFeed)
	}
	return feeds
}