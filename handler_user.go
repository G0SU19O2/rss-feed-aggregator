package main

import "fmt"

func handlerLogin(s *state, cmd command) error {
	argLen := len(cmd.Args)
	if argLen == 0 {
		return fmt.Errorf("missing argument")
	}
	username := cmd.Args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}
	fmt.Printf("username has been set to %s", username)
	return nil
}
