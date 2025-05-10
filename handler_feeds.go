package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("don't support argument")
	}
	feeds, err := s.db.GetFeeds(context.Background())
	if len(feeds) == 0 {
		fmt.Println("No feeds found.")
		return nil
	}
	if err != nil {
		return fmt.Errorf("couldn't list feeds: %w", err)
	}
	for _, feed := range feeds {
		username, err := s.db.GetUserName(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't get user: %w", err)
		}
		fmt.Println("=====================================")
		fmt.Println(username)
		fmt.Println(feed.Name)
		fmt.Println(feed.Url)
	}
	fmt.Println("Get feeds successfully!")
	return nil
}
