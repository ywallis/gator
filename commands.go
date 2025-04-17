package main 

import (
	"fmt"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	commandMap map[string]func(*State, command) error
}

func (c commands) register(name string, f func(*State, command) error) {
	c.commandMap[name] = f
}

func (c commands) Run(s *State, cmd command) error {

	if command, ok := c.commandMap[cmd.Name]; !ok {
		return fmt.Errorf("command not registered")
	} else {
		if err := command(s, cmd); err != nil{
			return err
		}
	}

	return nil
}

