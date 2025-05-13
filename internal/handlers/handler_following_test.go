package handlers

import (
	"context"
	"testing"
	"time"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/testutil"
)

func TestHandlerFollowingWithArgs(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	defer cleanup()
	cmd := cli.Command{Name: "following", Args: []string{"unexpected"}}
	if err := HandlerFollowing(state, cmd, database.User{}); err == nil {
		t.Error("Expected error for passing arguments, got nil")
	}
}

func TestHandlerFollowingNoFollows(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	defer cleanup()
	username := "test_following_nofollows"
	user := testutil.CreateTestUser(t, state.Db, username)
	defer state.Db.DeleteUser(context.Background(), username)
	if err := state.Cfg.SetUser(username); err != nil {
		t.Fatal("Failed to set user")
	}
	cmd := cli.Command{Name: "following", Args: []string{}}
	if err := HandlerFollowing(state, cmd, user); err != nil {
		t.Errorf("Expected no error for no follows, got: %v", err)
	}
}

func TestHandlerFollowingWithFollows(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	defer cleanup()
	username := "test_following_withfollows"
	testutil.CreateTestUser(t, state.Db, username)
	defer state.Db.DeleteUser(context.Background(), username)
	if err := state.Cfg.SetUser(username); err != nil {
		t.Fatal("Failed to set user")
	}
	feedID := "feed-id-456"
	feedURL := "http://testfeed2.com/rss"
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
		Name:      "Feed2",
		Url:       feedURL,
		UserID:    user.ID,
	})
	if err != nil {
		t.Fatalf("Failed to create test feed: %v", err)
	}
	if err := state.Db.CreateFeedFollow(context.Background(), struct {
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
	cmd := cli.Command{Name: "following", Args: []string{}}
	if err := HandlerFollowing(state, cmd, user); err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
}
