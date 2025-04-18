package main

import (
	"fmt"
	"time"
)

func HandlerAgg(s *State, cmd command) error {

	if len(cmd.Args) != 1 {
		return fmt.Errorf("Agg expects a single duration argument")
	}

	frequency, err := time.ParseDuration(cmd.Args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Collecting feeds every %s\n", frequency)
	ticker := time.NewTicker(frequency)
	for ; ; <- ticker.C {
		scrapeFeeds(s)
	}
	// feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	// if err != nil {
	// 	return err
	// }
	// fmt.Println(feed)
	// return nil
}
