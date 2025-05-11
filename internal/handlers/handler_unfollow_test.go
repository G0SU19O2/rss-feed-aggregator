package handlers

import (
	"context"
	"testing"
	"time"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
)

func TestHandlerUnfollowNoArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	user := createTestUser(t, state.Db, "test_unfollow_no_args")
	defer state.Db.DeleteUser(context.Background(), user.Name)
	cmd := cli.Command{Name: "unfollow", Args: []string{}}
	err := HandlerUnfollow(state, cmd, user)
	if err == nil {
		t.Error("Expected error for missing argument, got nil")
	}
}

func TestHandlerUnfollowTooManyArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	user := createTestUser(t, state.Db, "test_unfollow_many_args")
	defer state.Db.DeleteUser(context.Background(), user.Name)
	cmd := cli.Command{Name: "unfollow", Args: []string{"url1", "url2"}}
	err := HandlerUnfollow(state, cmd, user)
	if err == nil {
		t.Error("Expected error for too many arguments, got nil")
	}
}

func TestHandlerUnfollowFeedNotFound(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	user := createTestUser(t, state.Db, "test_unfollow_feed_notfound")
	defer state.Db.DeleteUser(context.Background(), user.Name)
	cmd := cli.Command{Name: "unfollow", Args: []string{"http://notfound.com/rss"}}
	err := HandlerUnfollow(state, cmd, user)
	if err == nil {
		t.Error("Expected error for feed not found, got nil")
	}
}

func TestHandlerUnfollowSuccess(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	username := "test_unfollow_success"
	user := createTestUser(t, state.Db, username)
	defer state.Db.DeleteUser(context.Background(), username)

	feedID := "feed-id-unfollow"
	feedURL := "http://testunfollow.com/rss"
	_, err := state.Db.CreateFeed(context.Background(), struct {
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
		Name:      "FeedUnfollow",
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		t.Fatalf("Failed to create test feed: %v", err)
	}
	err = state.Db.CreateFeedFollow(context.Background(), struct {
		ID        string
		CreatedAt time.Time
		UpdatedAt time.Time
		UserID    string
		FeedID    string
	}{
		ID:        "follow-id-unfollow",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feedID,
	})
	if err != nil {
		t.Fatalf("Failed to follow feed: %v", err)
	}

	cmd := cli.Command{Name: "unfollow", Args: []string{feedURL}}
	err = HandlerUnfollow(state, cmd, user)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}
