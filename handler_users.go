package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("don't support argument")
	}
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}
	for _, user := range users {
		username := user.Name
		if username != s.cfg.CurrentUserName {
			fmt.Printf("* %s\n", username)
		} else {
			fmt.Printf("* %s (current)\n", username)
		}
	}
	fmt.Println("Get users successfully!")
	return nil
}
