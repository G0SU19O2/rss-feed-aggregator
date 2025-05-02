package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("don't support argument")
	}
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("Channel Title: %s\n", feed.Channel.Title)
	fmt.Printf("Chanel Description: %s\n", feed.Channel.Description)
	for _, item := range feed.Channel.Item {
		fmt.Println("----------------------")
		fmt.Printf("Title: %s\n", item.Title)
		fmt.Printf("Link: %s\n", item.Link)
		fmt.Printf("Publish Date: %s\n", item.PubDate)
		fmt.Printf("Description: %s\n", item.Description)
	}
	return nil
}
