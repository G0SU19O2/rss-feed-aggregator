package main

import (
	"context"
	"fmt"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
)

func handlerReset(s *cli.State, cmd cli.Command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("don't support argument")
	}
	if err := s.Db.DeleteUsers(context.Background()); err != nil {
		return fmt.Errorf("couldn't delete users: %w", err)
	}
	fmt.Println("Database reset successfully!")
	return nil
}
