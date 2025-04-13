package main

import (
	"fmt"
	"os"

	"github.com/ywallis/gator/internal/cli"
	"github.com/ywallis/gator/internal/config"
)

func main() {
	var conf config.Config = config.ReadConfig()

	state := cli.State{Cfg: &conf}
	commands := cli.Commands{CommandMap: map[string]func(*cli.State, cli.Command) error{}}
	commands.Register("login", cli.HandlerLogin)
	userArgs := os.Args
	if len(userArgs) < 2 {
		fmt.Println("Not enough arguments in call.")
		return
	}
	currentCommand := cli.Command{Name: userArgs[1], Args: userArgs[2:]}
	if err := commands.Run(&state, currentCommand); err != nil {
		fmt.Println(err)
	}
}

