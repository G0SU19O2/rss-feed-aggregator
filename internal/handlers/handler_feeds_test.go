package handlers

import (
	"testing"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
)

func TestHandlerUsersFailWithArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := cli.Command{Name: "agg", Args: []string{"dummy"}}
	defer cleanup()
	if err := HandlerFeeds(state, cmd); err == nil {
		t.Error("Expected error because not enough arguments, got successful")
	}
}

func TestHandlerFeeds(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := cli.Command{Name: "feeds", Args: []string{}}
	defer cleanup()
	if err := HandlerFeeds(state, cmd); err != nil {
		t.Error("Fail to get feeds")
	}
}
