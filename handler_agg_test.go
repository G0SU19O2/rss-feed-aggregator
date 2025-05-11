package main

import (
	"testing"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
)

func TestHandlerAggFailWithArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := cli.Command{Name: "agg", Args: []string{"dummy"}}
	defer cleanup()
	if err := handlerAgg(state, cmd); err == nil {
		t.Error("Expected error because command have args, got successful")
	}
}

func TestHandlerAgg(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := cli.Command{Name: "agg", Args: []string{}}
	defer cleanup()
	if err := handlerAgg(state, cmd); err != nil {
		t.Errorf("Expected successful got error: %v", err)
	}
}
