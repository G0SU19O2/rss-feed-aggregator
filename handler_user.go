package main

import (
	"context"
	"fmt"
	"time"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *cli.State, cmd cli.Command) error {
	argLen := len(cmd.Args)
	if argLen == 0 {
		return fmt.Errorf("missing argument")
	}
	username := cmd.Args[0]
	user, err := s.Db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	err = s.Cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *cli.State, cmd cli.Command) error {
	argLen := len(cmd.Args)
	if argLen == 0 {
		return fmt.Errorf("missing argument")
	}
	username := cmd.Args[0]
	_, err := s.Db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New().String(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: username})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	err = s.Cfg.SetUser(username)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Printf("username %s has been created", username)
	return nil
}
