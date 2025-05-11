package handlers

import (
	"context"
	"fmt"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
)

func HandlerFollowing(s *cli.State, cmd cli.Command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("command does not support argument")
	}
	result, err := s.Db.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return fmt.Errorf("can't find following feed for login user: %v", err)
	}
	if len(result) == 0 {
		fmt.Println("No feed follows found for this user.")
		return nil
	}
	fmt.Println("Feed name----------------------")
	for _, item := range result {
		fmt.Printf("%s\n", item.FeedName)
	}
	return nil
}
