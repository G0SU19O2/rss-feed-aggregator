package handlers

import (
	"context"
	"fmt"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
)

func HandlerUnfollow(s *cli.State, cmd cli.Command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("only support one argument")
	}
	feedURL := cmd.Args[0]
	feed, err := s.Db.GetFeedByURL(context.Background(), feedURL)
	if err != nil {
		return fmt.Errorf("couldn't get feed: %w", err)
	}
	err = s.Db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed.ID})
	if err != nil {
		return fmt.Errorf("couldn't unfollow feed: %w", err)
	}
	fmt.Println("Unfollow feed successfully")
	return nil
}
