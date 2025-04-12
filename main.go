package main

import "github.com/ywallis/gator/internal/config"
import "fmt"

func main() {
	var conf config.Config = config.ReadConfig()

	fmt.Println(conf)
	conf.SetUser("Tanja")
	fmt.Println(conf)
}
