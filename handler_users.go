package main

import (
	"context"
	"fmt"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
)

func handlerUsers(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("don't support argument")
	}
	users, err := s.Db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't list users: %w", err)
	}
	for _, user := range users {
		username := user.Name
		if username != s.Cfg.CurrentUserName {
			fmt.Printf("* %s\n", username)
		} else {
			fmt.Printf("* %s (current)\n", username)
		}
	}
	fmt.Println("Get users successfully!")
	return nil
}
