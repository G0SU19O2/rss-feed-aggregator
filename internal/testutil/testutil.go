package testutil

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/cli"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/config"
	"github.com/G0SU19O2/rss-feed-aggregator/internal/database"
	"github.com/google/uuid"
)

func SetupTestDB(t *testing.T) (*cli.State, func()) {
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

func CreateTestUser(t *testing.T, db *database.Queries, username string) database.User {
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

func GetCurrentUser(t *testing.T, s *cli.State) database.User {
	user, err := s.Db.GetUser(context.Background(), s.Cfg.CurrentUserName)
	if err != nil {
		t.Fatal("Failed to get current user")
	}
	return user
}
