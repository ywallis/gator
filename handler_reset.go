package main

import (
	"context"
	"fmt"
)

func HandlerReset(s *State, cmd command) error {

	err := s.db.DeleteAllUsers(context.Background())
	if err != nil {
		return fmt.Errorf("Could not delete all users: %s", err)
	}

	fmt.Println("All users deleted from db")

	return nil
}
