package main

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/ywallis/gator/internal/database"
)

func HandlerRegister(s *State, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Register function requires an argument.")
	}

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.Args[0]})
	if err != nil {
		return fmt.Errorf("Error creating user. Likely already exists. Err: %s", err)
	}

	s.cfg.SetUser(user.Name)
	fmt.Printf("User %s was created, full data: %s\n", user.Name, user)

	return nil
}
