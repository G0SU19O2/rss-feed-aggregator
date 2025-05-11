package cli

import "fmt"

type Command struct {
	Name string
	Args []string
}
type Commands struct {
	RegisteredCommands map[string]func(*State, Command) error
}

func (c *Commands) Run(s *State, cmd Command) error {
	f, ok := c.RegisteredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("command doesn't exist")
	}
	if err := f(s, cmd); err != nil {
		return err
	}
	return nil
}

func (c *Commands) Register(Name string, f func(*State, Command) error) {
	c.RegisteredCommands[Name] = f
}
