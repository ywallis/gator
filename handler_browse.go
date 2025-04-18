package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ywallis/gator/internal/database"
)

func HandlerBrowse(s *State, cmd command, user database.User) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("The browse command takes one argument, the amount of posts wanted.")
	}
	limit, err := strconv.ParseInt(cmd.Args[0], 0, 64)
	if err != nil {
		return err
	}
	
	queryParams := database.GetPostsForUserParams{
		ID:    user.ID,
		Limit: int32(limit),
	}
	userPosts, err := s.db.GetPostsForUser(context.Background(), queryParams)

	for _, post := range userPosts {
		fmt.Println(post.Title)
		fmt.Println(post.Description.String)
	}

	return nil
}
