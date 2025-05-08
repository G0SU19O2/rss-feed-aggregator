package main

import (
	"context"
	"testing"
)

func TestHandlerAddFeedFailWithArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := command{Name: "agg", Args: []string{"dummy"}}
	defer cleanup()
	if err := handlerAddFeed(state, cmd); err == nil {
		t.Error("Expected error because not enough arguments, got successful")
	}
}

func TestHandlerAddFeed(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()
	username := "test"
	createTestUser(t, state.db, username)
	defer state.db.DeleteUser(context.Background(), username)
	if err := state.cfg.SetUser(username); err != nil {
		t.Error("Failed to set user")
	}
	cmd := command{Name: "agg", Args: []string{"feedName", "feedURL"}}
	if err := handlerAddFeed(state, cmd); err != nil {
		t.Errorf("Expected successful create feed, got error: %v", err)
	}
}
