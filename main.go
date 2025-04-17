package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/ywallis/gator/internal/config"
	"github.com/ywallis/gator/internal/database"
)

type State struct {
	cfg *config.Config
	db *database.Queries
}


func main() {

	var conf config.Config
	conf, err := config.ReadConfig()
	if err != nil {
		fmt.Printf("Error reading config: %s", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", conf.DbUrl)
	if err != nil {
		fmt.Printf("Error connecting to db: %s", err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	state := State{cfg: &conf, db: dbQueries}
	commands := commands{commandMap: map[string]func(*State, command) error{}}
	commands.register("login", HandlerLogin)
	commands.register("register", HandlerRegister)
	commands.register("reset", HandlerReset)
	commands.register("users", HandlerGetUsers)
	commands.register("agg", HandlerAgg)
	commands.register("addfeed", middlewareLoggedIn(HandlerAddFeed))
	commands.register("feeds", HandlerFeeds)
	commands.register("follow", middlewareLoggedIn(HandlerFollowFeed))
	commands.register("following", middlewareLoggedIn(HandlerFollowing))
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
