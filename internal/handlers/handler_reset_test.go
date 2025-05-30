package handlers

import (
	"testing"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/testutil"
)

func TestResetUsers(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	cmd := cli.Command{Name: "reset", Args: []string{}}
	defer cleanup()
	if err := HandlerReset(state, cmd); err != nil {
		t.Error("Fail to reset users")
	}
}

func TestResetUsersFailWithArgs(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	cmd := cli.Command{Name: "reset", Args: []string{"dummy"}}
	defer cleanup()
	if err := HandlerReset(state, cmd); err == nil {
		t.Error("Expected error because command have args, got successful")
	}
}
