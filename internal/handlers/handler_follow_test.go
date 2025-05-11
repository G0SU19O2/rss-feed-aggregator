package handlers

import (
	"context"
	"testing"
	"time"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
)

func TestHandlerFollowFailWithArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	cmd := cli.Command{Name: "follow", Args: []string{}}
	if err := HandlerFollow(state, cmd); err == nil {
		t.Error("Expected error because no argument, got successful")
	}
	cmd = cli.Command{Name: "follow", Args: []string{"a", "b"}}
	if err := HandlerFollow(state, cmd); err == nil {
		t.Error("Expected error because too many arguments, got successful")
	}
}

func TestHandlerFollowFeedNotFound(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	username := "test_follow_feed_not_found"
	createTestUser(t, state.Db, username)
	defer state.Db.DeleteUser(context.Background(), username)
	if err := state.Cfg.SetUser(username); err != nil {
		t.Fatal("Failed to set user")
	}
	cmd := cli.Command{Name: "follow", Args: []string{"nonexistent_feed_url"}}
	if err := HandlerFollow(state, cmd); err == nil {
		t.Error("Expected error for non-existent feed, got nil")
	}
}

func TestHandlerFollowUserNotFound(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	cmd := cli.Command{Name: "follow", Args: []string{"some_feed_url"}}
	if err := HandlerFollow(state, cmd); err == nil {
		t.Error("Expected error for non-existent user, got nil")
	}
}

func TestHandlerFollowSuccess(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	username := "test_follow_success"
	createTestUser(t, state.Db, username)
	defer state.Db.DeleteUser(context.Background(), username)
	if err := state.Cfg.SetUser(username); err != nil {
		t.Fatal("Failed to set user")
	}
	feedID := "feed-id-123"
	feedURL := "http://testfeed.com/rss"
	user, err := state.Db.GetUser(context.Background(), username)
	if err != nil {
		t.Fatal("Failed to get test user")
	}
	_, err = state.Db.CreateFeed(context.Background(), struct {
		ID        string
		CreatedAt time.Time
		UpdatedAt time.Time
		Name      string
		Url       string
		UserID    string
	}{
		ID:        feedID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Test Feed",
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		t.Fatalf("Failed to create test feed: %v", err)
	}
	cmd := cli.Command{Name: "follow", Args: []string{feedURL}}
	if err := HandlerFollow(state, cmd); err != nil {
		t.Errorf("Expected successful follow, got error: %v", err)
	}
}
