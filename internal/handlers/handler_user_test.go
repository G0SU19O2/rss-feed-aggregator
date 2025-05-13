package handlers

import (
	"context"
	"testing"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/testutil"
	_ "github.com/go-sql-driver/mysql"
)

func TestLoginWithValidUser(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	username := "test_login"
	testutil.CreateTestUser(t, state.Db, username)
	defer state.Db.DeleteUser(context.Background(), username)

	cmd := cli.Command{Name: "login", Args: []string{username}}
	err := HandlerLogin(state, cmd)

	if err != nil {
		t.Errorf("Expected successful login, got error: %v", err)
	}

	if state.Cfg.CurrentUserName != username {
		t.Errorf("Want username: %s, got %s", username, state.Cfg.CurrentUserName)
	}
}

func TestLoginWithNonExistentUser(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	cmd := cli.Command{Name: "login", Args: []string{"nonexistent"}}
	err := HandlerLogin(state, cmd)

	if err == nil {
		t.Error("Expected error for non-existent user, got nil")
	}
}

func TestRegisterNewUser(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	username := "test_register"
	cmd := cli.Command{Name: "register", Args: []string{username}}

	err := HandlerRegister(state, cmd)
	if err != nil {
		t.Errorf("Expected successful registration, got error: %v", err)
	}

	defer state.Db.DeleteUser(context.Background(), username)

	// Verify user was created
	user, err := state.Db.GetUser(context.Background(), username)
	if err != nil {
		t.Errorf("Failed to get created user: %v", err)
	}
	if user.Name != username {
		t.Errorf("Want username: %s, got %s", username, user.Name)
	}
}

func TestRegisterDuplicateUser(t *testing.T) {
	state, cleanup := testutil.SetupTestDB(t)
	defer cleanup()

	username := "duplicate_user"
	testutil.CreateTestUser(t, state.Db, username)
	defer state.Db.DeleteUser(context.Background(), username)

	cmd := cli.Command{Name: "register", Args: []string{username}}
	err := HandlerRegister(state, cmd)

	if err == nil {
		t.Error("Expected error for duplicate username, got nil")
	}
}
