package handlers

import (
	"testing"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
)

func TestResetUsers(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := cli.Command{Name: "reset", Args: []string{}}
	defer cleanup()
	if err := HandlerReset(state, cmd); err != nil {
		t.Error("Fail to reset users")
	}
}

func TestResetUsersFailWithArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := cli.Command{Name: "reset", Args: []string{"dummy"}}
	defer cleanup()
	if err := HandlerReset(state, cmd); err == nil {
		t.Error("Expected error because command have args, got successful")
	}
}
