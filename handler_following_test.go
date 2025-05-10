package main

import (
	"context"
	"testing"
	"time"
)

func TestHandlerFollowingWithArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	cmd := command{Name: "following", Args: []string{"unexpected"}}
	if err := handlerFollowing(state, cmd); err == nil {
		t.Error("Expected error for passing arguments, got nil")
	}
}

func TestHandlerFollowingUserNotFound(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	cmd := command{Name: "following", Args: []string{}}
	if err := handlerFollowing(state, cmd); err == nil {
		t.Error("Expected error for non-existent user, got nil")
	}
}

func TestHandlerFollowingNoFollows(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	username := "test_following_nofollows"
	createTestUser(t, state.db, username)
	defer state.db.DeleteUser(context.Background(), username)
	if err := state.cfg.SetUser(username); err != nil {
		t.Fatal("Failed to set user")
	}
	cmd := command{Name: "following", Args: []string{}}
	if err := handlerFollowing(state, cmd); err != nil {
		t.Errorf("Expected no error for no follows, got: %v", err)
	}
}

func TestHandlerFollowingWithFollows(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	username := "test_following_withfollows"
	createTestUser(t, state.db, username)
	defer state.db.DeleteUser(context.Background(), username)
	if err := state.cfg.SetUser(username); err != nil {
		t.Fatal("Failed to set user")
	}
	feedID := "feed-id-456"
	feedURL := "http://testfeed2.com/rss"
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
		Name:      "Feed2",
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		t.Fatalf("Failed to create test feed: %v", err)
	}
	if err := state.db.CreateFeedFollow(context.Background(), struct {
		ID        string
		CreatedAt time.Time
		UpdatedAt time.Time
		UserID    string
		FeedID    string
	}{
		ID:        "follow-id-1",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedID,
	}); err != nil {
		t.Fatalf("Failed to follow feed: %v", err)
	}
	cmd := command{Name: "following", Args: []string{}}
	if err := handlerFollowing(state, cmd); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}
