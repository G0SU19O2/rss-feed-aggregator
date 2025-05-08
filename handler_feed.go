package main

import (
	"context"
	"database/sql"
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
	_, err = s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      sql.NullString{String: feedName, Valid: true},
		Url:       sql.NullString{String: feedURL, Valid: true},
		UserID:    sql.NullString{String: user.ID, Valid: true},
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}
	return nil
}
