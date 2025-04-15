package main

import (
	"context"
	"fmt"
)

func HandlerLogin(s *State, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Login function requires an argument.")
	}

	user, err := s.db.GetUser(context.Background(), cmd.Args[0]) 
	if err != nil {
		return fmt.Errorf("Username does not exist in database!")

	}

	if err := s.cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Printf("User %s has been set.\n", cmd.Args[0])
	return nil
}
