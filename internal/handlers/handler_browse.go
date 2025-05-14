package handlers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
)

func HandlerBrowse(s *cli.State, cmd cli.Command, user database.User) error {
	limit := 2

	if len(cmd.Args) >= 1 {
		parsedLimit, err := strconv.Atoi(cmd.Args[0])
		if err == nil && parsedLimit > 0 {
			limit = parsedLimit
		} else if err != nil {
			return fmt.Errorf("invalid limit: %s, using default limit of %d", cmd.Args[0], limit)
		}
	}

	posts, err := s.Db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("error fetching posts: %w", err)
	}

	// Display posts
	fmt.Printf("Recent posts (limit: %d):\n\n", limit)
	if len(posts) == 0 {
		fmt.Println("No posts found. Follow some feeds to see posts.")
		return nil
	}

	for i, post := range posts {
		fmt.Printf("%d. %s\n", i+1, post.Title)
		fmt.Printf("   %s\n", post.Url)
		fmt.Printf("   Published: %s\n", post.PublishedAt.Format("Jan 2, 2006"))
		if post.Description != "" {
			// Display a limited preview of the description
			desc := post.Description
			if len(desc) > 100 {
				desc = desc[:100] + "..."
			}
			fmt.Printf("   %s\n", desc)
		}
		fmt.Println()
	}

	return nil
}
