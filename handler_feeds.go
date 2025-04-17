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
func HandlerAddFeed(s *State, cmd command, user database.User) error {

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

	fmt.Printf("Created feed with name: %s and url: %s\n", record.Name, record.Url)

	queryParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    record.ID,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), queryParams)

	fmt.Printf("User %s started following feed %s\n", feedFollow.UserName, feedFollow.FeedName)

	return nil
}

func HandlerFollowFeed(s *State, cmd command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("Command expects one url argument.")
	}

	url := cmd.Args[0]
	feed, err := s.db.GetFeedFromUrl(context.Background(), url)
	if err != nil {
		return err
	}

	queryParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	feedFollow, err := s.db.CreateFeedFollow(context.Background(), queryParams)

	fmt.Printf("User %s started following feed %s\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func HandlerFollowing(s *State, cmd command, user database.User) error {


	if len(cmd.Args) != 0 {
		return fmt.Errorf("Command expects no arguments.")
	}

	allFollowed, err := s.db.GetFeedFollowForUser(context.Background(), user.ID)
	if err != nil {
		return err
	}

	fmt.Printf("Current user (%s) is following:\n", user.Name)

	for _, row := range allFollowed {
		fmt.Println(row.FeedName)
	}
	return nil
}

func HandlerUnfollow(s *State, cmd command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("Command expects one url argument")
	}
	
	url := cmd.Args[0]

	feed, err := s.db.GetFeedFromUrl(context.Background(), url)
	if err != nil {
		return err
	}
	queryParams := database.UnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	if err := s.db.Unfollow(context.Background(), queryParams); err != nil {
		return err
	}

	fmt.Printf("User %s unfollowed feed %s\n", user.Name, feed.Name)

	return nil
}
