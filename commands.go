package main

import "fmt"

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	val, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("command not found: %s", cmd)
	}
	return val(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f

}
