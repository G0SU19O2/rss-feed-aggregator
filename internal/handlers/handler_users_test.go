package handlers

import (
	"testing"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/testutil"
)

func TestGetUsers(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	cmd := cli.Command{Name: "users", Args: []string{}}
	defer cleanup()
	if err := HandlerReset(state, cmd); err != nil {
		t.Error("Fail to get users")
	}
}

func TestGetUsersFailWithArgs(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	cmd := cli.Command{Name: "users", Args: []string{"dummy"}}
	defer cleanup()
	if err := HandlerReset(state, cmd); err == nil {
		t.Error("Expected error because command have args, got successful")
	}
}
