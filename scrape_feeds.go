package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
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
		// publishedTime := time.Now()
		// fmt.Printf("DEBUG published %s\n", item.PubDate)
		publishedTime, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			return err
		}
		postQueryParams := database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:       item.Title,
			Url:         item.Link,
			Description: sql.NullString{String: item.Description, Valid: true},
			PublishedAt: sql.NullTime{Time: publishedTime, Valid: true},
			FeedID:      feed.ID,
		}
		s.db.CreatePost(context.Background(), postQueryParams)
		fmt.Printf("Post %s saved to DB\n", item.Title)
	}

	return nil
}
