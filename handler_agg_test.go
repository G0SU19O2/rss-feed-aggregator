package main

import "testing"


func TestHandlerAggFailWithArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := command{Name: "agg", Args: []string{"dummy"}}
	defer cleanup()
	if err := handlerAgg(state, cmd); err == nil {
		t.Error("Expected error because command have args, got successful")
	}
}

func TestHandlerAgg(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := command{Name: "agg", Args: []string{}}
	defer cleanup()
	if err := handlerAgg(state, cmd); err != nil {
		t.Errorf("Expected successful got error: %v", err)
	}
}