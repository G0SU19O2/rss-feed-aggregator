package handlers

import (
	"context"
	"testing"
	"time"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/testutil"
	"github.com/google/uuid"
)

func TestHandlerBrowseFailWithToManyArgs(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	cmd := cli.Command{Name: "browse", Args: []string{"2", "3"}}
	defer cleanup()
	if err := HandlerAgg(state, cmd); err == nil {
		t.Error("Expected error because command only support zero or one arg")
	}
}

func TestHandlerBrowse(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	defer cleanup()
	username := "test"
	testutil.CreateTestUser(t, state.Db, username)
	defer state.Db.DeleteUser(context.Background(), username)
	if err := state.Cfg.SetUser(username); err != nil {
		t.Error("Failed to set user")
	}
	feedID := uuid.NewString()
	user := testutil.GetCurrentUser(t, state)
	_, err := state.Db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        feedID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      "Name",
		Url:       "Url",
		UserID:    user.ID,
	})
	if err != nil {
		t.Errorf("couldn't create feed: %v", err)
	}
	defer state.Db.DeleteFeed(context.Background(), feedID)
	state.Db.CreatePost(context.Background(), database.CreatePostParams{
		ID:          uuid.NewString(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       "Title",
		Url:         "URL",
		Description: "Description",
		PublishedAt: time.Time{},
		FeedID:      feedID,
	})
	cmd := cli.Command{Name: "browse", Args: []string{"5"}}
	if err := HandlerBrowse(state, cmd, user); err != nil {
		t.Errorf("Failed to browse %v", err)
	}
}
