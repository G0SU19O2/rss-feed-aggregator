package handlers

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/config"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
	_ "github.com/go-sql-driver/mysql"
)

func setupTestDB(t *testing.T) (*cli.State, func()) {
	t.Helper()
	cfg, err := config.Read()
	if err != nil {
		t.Fatalf("Failed to read config: %v", err)
	}

	db, err := sql.Open("mysql", cfg.Db_url)
	if err != nil {
		t.Fatalf("Failed to open DB connection: %v", err)
	}

	dbQueries := database.New(db)
	programState := &cli.State{
		Cfg: &cfg,
		Db:  dbQueries,
	}

	cleanup := func() {
		db.Close()
	}

	return programState, cleanup
}

func createTestUser(t *testing.T, db *database.Queries, username string) database.User {
	t.Helper()
	id := uuid.New().String()
	createdAt := time.Now()
	updatedAt := time.Now()
	name := username
	_, err := db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        id,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		Name:      name,
	})
	if err != nil {
		t.Fatalf("Failed to create test user: %v", err)
	}
	return database.User{ID: id, CreatedAt: createdAt, UpdatedAt: createdAt, Name: name}
}

func TestLoginWithValidUser(t *testing.T) {
	state, cleanup := setupTestDB(t)
	defer cleanup()

	username := "test_login"
	createTestUser(t, state.Db, username)
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
	state, cleanup := setupTestDB(t)
	defer cleanup()

	cmd := cli.Command{Name: "login", Args: []string{"nonexistent"}}
	err := HandlerLogin(state, cmd)

	if err == nil {
		t.Error("Expected error for non-existent user, got nil")
	}
}

func TestRegisterNewUser(t *testing.T) {
	state, cleanup := setupTestDB(t)
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
	state, cleanup := setupTestDB(t)
	defer cleanup()

	username := "duplicate_user"
	createTestUser(t, state.Db, username)
	defer state.Db.DeleteUser(context.Background(), username)

	cmd := cli.Command{Name: "register", Args: []string{username}}
	err := HandlerRegister(state, cmd)

	if err == nil {
		t.Error("Expected error for duplicate username, got nil")
	}
}
