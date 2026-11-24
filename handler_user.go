package main

import (
	"errors"
	"fmt"
	"os"
)

// the function signature of all command handlers
func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("the login handler expects a single argument, the username")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		return err

	}
	fmt.Println("username has been set")
	return nil
}
