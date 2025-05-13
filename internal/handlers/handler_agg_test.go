package handlers

import (
	"testing"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/testutil"
)

func TestHandlerAggFailWithArgs(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	cmd := cli.Command{Name: "agg", Args: []string{"dummy"}}
	defer cleanup()
	if err := HandlerAgg(state, cmd); err == nil {
		t.Error("Expected error because command must have time duration arg")
	}
}
