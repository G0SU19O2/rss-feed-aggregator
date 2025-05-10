package main

import (
	"context"
	"fmt"
	"time"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.Args) < 2 {
		return fmt.Errorf("missing arguments")
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("can't find current login user: %v", err)
	}
	feedName := cmd.Args[0]
	feedURL := cmd.Args[1]
	feedID := uuid.New().String()
	_, err = s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        feedID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      feedName,
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	fmt.Println("Feed created successfully")
	err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Println("Feed followed successfully")
	return nil
}
