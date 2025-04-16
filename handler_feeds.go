package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ywallis/gator/internal/database"
)

func HandlerFeeds(s *State, cmd command) error {

	feeds, err := s.db.GetFeedsWithUser(context.Background())
	if err != nil {
		return err
	}

	for i, feed := range feeds {
		fmt.Printf("%d: %s, %s (added by %s)\n", i, feed.Name, feed.Url, feed.Name_2.String)
	}
	return nil
}
func HandlerAddFeed(s *State, cmd command) error {

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}
	if len(cmd.Args) != 2 {
		return fmt.Errorf("Add feed should take 2 arguments, name and url")
	}
	name := cmd.Args[0]
	url := cmd.Args[1]

	inputParameters := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	record, err := s.db.CreateFeed(context.Background(), inputParameters)
	if err != nil {
		return err
	}

	fmt.Printf("Created feed with name: %s and url: %s", record.Name, record.Url)

	return nil
}
