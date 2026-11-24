package main

import (
	"fmt"
	"github.com/GritandGo/blog-aggregator/internal/config"
	"os"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	s := state{
		cfg: &cfg,
	}
}

type state struct {
	cfg *config.Config
}

commandRegistry := commands{
	handlers: make(map[string]func(*state, command) error),
}

commandRegistry.register("login", handlerLogin)