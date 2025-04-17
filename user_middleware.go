package main

import (
	"context"

	"github.com/ywallis/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *State, cmd command, user database.User) error) func(s *State, cmd command) error {

	return func(s *State, cmd command) error {

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}

}
