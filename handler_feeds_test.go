package main

import (
	"testing"
)

func TestHandlerUsersFailWithArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := command{Name: "agg", Args: []string{"dummy"}}
	defer cleanup()
	if err := handlerFeeds(state, cmd); err == nil {
		t.Error("Expected error because not enough arguments, got successful")
	}
}

func TestHandlerFeeds(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := command{Name: "feeds", Args: []string{}}
	defer cleanup()
	if err := handlerFeeds(state, cmd); err != nil {
		t.Error("Fail to get feeds")
	}
}
