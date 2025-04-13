package main

import "fmt"

func HandlerLogin(s *State, cmd command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Login function requires an argument.")
	}

	if err := s.Cfg.SetUser(cmd.Args[0]); err != nil {
		return err
	}

	fmt.Printf("User %s has been set.\n", cmd.Args[0])
	return nil
}
