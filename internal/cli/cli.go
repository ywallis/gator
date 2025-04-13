package cli

import (
	"fmt"

	"github.com/ywallis/gator/internal/config"
)

type State struct {
	Cfg *config.Config
}

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CommandMap map[string]func(*State, Command) error
}

func (c Commands) Register(name string, f func(*State, Command) error) {
	c.CommandMap[name] = f
}

func (c Commands) Run(s *State, cmd Command) error {

	if command, ok := c.CommandMap[cmd.Name]; !ok {
		return fmt.Errorf("Command not registered")
	} else {
		command(s, cmd)
	}

	return nil
}

func HandlerLogin(s *State, cmd Command) error {
	if len(cmd.Args) == 0 {
		return fmt.Errorf("Login function requires an argument.")
	}

	if err := s.Cfg.SetUser(cmd.Args[0]); err != nil {
		return err
	}

	fmt.Printf("User %s has been set.\n", cmd.Args[0])
	return nil
}
