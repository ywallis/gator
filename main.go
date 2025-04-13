package main

import (
	"fmt"
	"os"

	"github.com/ywallis/gator/internal/config"
)

type State struct {
	Cfg *config.Config
}

func main() {
	var conf config.Config
	conf, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Error reading config: %s", err)
		os.Exit(1)
	}

	state := State{Cfg: &conf}
	commands := commands{commandMap: map[string]func(*State, command) error{}}
	commands.Register("login", HandlerLogin)
	userArgs := os.Args
	if len(userArgs) < 2 {
		fmt.Println("Not enough arguments in call.")
		os.Exit(1)
	}
	currentCommand := command{Name: userArgs[1], Args: userArgs[2:]}
	if err := commands.Run(&state, currentCommand); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
