package handlers

import (
	"context"
	"fmt"
	"time"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func HandlerFollow(s *cli.State, cmd cli.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("only support one argument")
	}
	feedURL := cmd.Args[0]
	feed, err := s.Db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}
	err = s.Db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{ID: uuid.New().String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create feed follow: %w", err)
	}
	fmt.Println("Feed follow created")
	return nil
}
