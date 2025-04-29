package main

import "fmt"

type command struct {
	Name      string
	Args []string
}
type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("command doesn't exist")
	}
	if err := f(s, cmd); err != nil {
		return err
	}
	return nil
}

func (c *commands) register(Name string, f func(*state, command) error) {
	c.registeredCommands[Name] = f
}
