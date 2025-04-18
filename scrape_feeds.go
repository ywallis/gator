package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/ywallis/gator/internal/database"
)

func scrapeFeeds(s *State) error {

	feed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}
	queryParams := database.MarkFeedFetchedParams{
		UpdatedAt: time.Now(),
		FetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
		ID:        feed.ID,
	}
	if err := s.db.MarkFeedFetched(context.Background(), queryParams); err != nil {
		return err
	}

	res, err := fetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range res.Channel.Item {
		fmt.Println(item.Title)
	}

	return nil
}
