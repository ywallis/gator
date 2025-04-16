package main

import (
	"context"
	"fmt"
)

func HandlerAgg(s *State, cmd command) error {

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Println(feed)
	return nil
}
