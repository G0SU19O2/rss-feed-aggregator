package main

import "testing"

func TestResetUsers(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := command{Name: "reset", Args: []string{}}
	defer cleanup()
	if err := handlerReset(state, cmd); err != nil {
		t.Error("Fail to reset users")
	}
}

func TestResetUsersFailWithArgs(t *testing.T) {
	state, cleanup := setupTestDB(t)
	cmd := command{Name: "reset", Args: []string{"dummy"}}
	defer cleanup()
	if err := handlerReset(state, cmd); err == nil {
		t.Error("Expected error because command have args, got successful")
	}
}
