package main

import (
	"fmt"
	"github.com/GritandGo/blog-aggregator/internal/config"
	"os"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	s := state{
		cfg: &cfg,
	}

	commandRegistry := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	commandRegistry.register("login", handlerLogin)


	userInput := os.Args
	if len(userInput) < 2 {
		fmt.Fprintln(os.Stderr, "Please enter a command after program name")
		os.Exit(1)
	}

	c := command{
		name: userInput[1],
		args: userInput[2:],
	}

	err = commandRegistry.run(&s, c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}
