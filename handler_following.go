package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("command does not support argument")
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("can't find current login user: %v", err)
	}
	result, err := s.db.GetFeedFollowsForUser(context.Background(), user.Name)
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
