package main

import (
	"context"
	"testing"
	"time"
)

func TestHandlerFollowFailWithArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	cmd := command{Name: "follow", Args: []string{}}
	if err := handlerFollow(state, cmd); err == nil {
		t.Error("Expected error because no argument, got successful")
	}
	cmd = command{Name: "follow", Args: []string{"a", "b"}}
	if err := handlerFollow(state, cmd); err == nil {
		t.Error("Expected error because too many arguments, got successful")
	}
}

func TestHandlerFollowFeedNotFound(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	username := "test_follow_feed_not_found"
	createTestUser(t, state.db, username)
	defer state.db.DeleteUser(context.Background(), username)
	if err := state.cfg.SetUser(username); err != nil {
		t.Fatal("Failed to set user")
	}
	cmd := command{Name: "follow", Args: []string{"nonexistent_feed_url"}}
	if err := handlerFollow(state, cmd); err == nil {
		t.Error("Expected error for non-existent feed, got nil")
	}
}

func TestHandlerFollowUserNotFound(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	cmd := command{Name: "follow", Args: []string{"some_feed_url"}}
	if err := handlerFollow(state, cmd); err == nil {
		t.Error("Expected error for non-existent user, got nil")
	}
}

func TestHandlerFollowSuccess(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	username := "test_follow_success"
	createTestUser(t, state.db, username)
	defer state.db.DeleteUser(context.Background(), username)
	if err := state.cfg.SetUser(username); err != nil {
		t.Fatal("Failed to set user")
	}
	feedID := "feed-id-123"
	feedURL := "http://testfeed.com/rss"
	user, err := state.db.GetUser(context.Background(), username)
	if err != nil {
		t.Fatal("Failed to get test user")
	}
	_, err = state.db.CreateFeed(context.Background(), struct {
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
	cmd := command{Name: "follow", Args: []string{feedURL}}
	if err := handlerFollow(state, cmd); err != nil {
		t.Errorf("Expected successful follow, got error: %v", err)
	}
}
