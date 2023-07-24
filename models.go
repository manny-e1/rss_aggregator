package main

import (
	"time"

	"github.com/google/uuid"
	"github.com/manny-e1/rss_aggregator/internal/database"
)

type User struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	ApiKey    string    `json:"apiKey"`
}

type Feed struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"userId"`
}

type FeedFollow struct {
	Id        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	FeedID    uuid.UUID `json:"feedId"`
	UserID    uuid.UUID `json:"userId"`
}

func dbUserToCustomUser(dbUser database.User) User {
	return User{
		Id:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
		ApiKey:    dbUser.ApiKey,
	}
}

func dbFeedToCustomFeed(dbFeed database.Feed) Feed {
	return Feed{
		Id:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		Name:      dbFeed.Name,
		Url:       dbFeed.Url,
		UserID:    dbFeed.UserID,
	}
}

func dbFeedsToCustomFeeds(dbFeeds []database.Feed) []Feed {
	feeds := []Feed{}
	for _, feed := range dbFeeds {
		feeds = append(feeds, dbFeedToCustomFeed(feed))
	}
	return feeds
}

func dbFeedFollowToCustomFeedFollow(dbFeed database.FeedFollow) FeedFollow {
	return FeedFollow{
		Id:        dbFeed.ID,
		CreatedAt: dbFeed.CreatedAt,
		UpdatedAt: dbFeed.UpdatedAt,
		FeedID:    dbFeed.FeedID,
		UserID:    dbFeed.UserID,
	}
}

func dbFeedFollowsToCustomFeedFollows(dbFeedFollows []database.FeedFollow) []FeedFollow {
	feedFollows := []FeedFollow{}
	for _, feedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, dbFeedFollowToCustomFeedFollow(feedFollow))
	}
	return feedFollows
}
