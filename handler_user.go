package main

import (
	"context"
	"fmt"
	"time"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerLogin(s *state, cmd command) error {
	argLen := len(cmd.Args)
	if argLen == 0 {
		return fmt.Errorf("missing argument")
	}
	username := cmd.Args[0]
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("couldn't find user: %w", err)
	}
	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}
	fmt.Println("User switched successfully!")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	argLen := len(cmd.Args)
	if argLen == 0 {
		return fmt.Errorf("missing argument")
	}
	username := cmd.Args[0]
	_, err := s.db.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New().String(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: username})
	if err != nil {
		return fmt.Errorf("couldn't create user: %w", err)
	}
	fmt.Printf("username %s has been created", username)
	return nil
}
