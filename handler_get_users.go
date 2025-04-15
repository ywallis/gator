package main

import (
	"context"
	"fmt"
)

func HandlerGetUsers(s *State, cmd command) error {

	users, err := s.db.FetchAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching users: %s", err)
	}
	for _, user := range users {
		if user.Name == s.cfg.CurrentUserName {
			fmt.Printf("- %s (current)\n", user.Name)
		} else {
			fmt.Printf("- %s\n", user.Name)
		}
	}
	return nil
}
