package main

import (
	"testing"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
)

func TestGetUsers(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := cli.Command{Name: "users", Args: []string{}}
	defer cleanup()
	if err := handlerReset(state, cmd); err != nil {
		t.Error("Fail to get users")
	}
}

func TestGetUsersFailWithArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := cli.Command{Name: "users", Args: []string{"dummy"}}
	defer cleanup()
	if err := handlerReset(state, cmd); err == nil {
		t.Error("Expected error because command have args, got successful")
	}
}
