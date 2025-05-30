package handlers

import (
	"context"
	"testing"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/testutil"
)

func TestHandlerAddFeedFailWithArgs(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	cmd := cli.Command{Name: "addfeed", Args: []string{"dummy"}}
	defer cleanup()
	if err := HandlerAddFeed(state, cmd, database.User{}); err == nil {
		t.Error("Expected error because not enough arguments, got successful")
	}
}

func TestHandlerAddFeed(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	defer cleanup()
	username := "test"
	testutil.CreateTestUser(t, state.Db, username)
	defer state.Db.DeleteUser(context.Background(), username)
	if err := state.Cfg.SetUser(username); err != nil {
		t.Error("Failed to set user")
	}
	cmd := cli.Command{Name: "addfeed", Args: []string{"feedName", "feedURL"}}
	user := testutil.GetCurrentUser(t, state)
	if err := HandlerAddFeed(state, cmd, user); err != nil {
		t.Errorf("Expected successful create feed, got error: %v", err)
	}
}
