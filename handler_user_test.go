package main

import (
	"testing"

	"github.com/G0SU19O2/rss-feed-aggregator/internal/config"
)

func TestHandlerLogin(t *testing.T) {
	tests := []struct {
		name     string
		command  string
		username string
		handler  func(*state, command) error
	}{
		{
			name:     "handler login successful",
			command:  "login",
			username: "G0SU19O2",
			handler:  handlerLogin,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, _ := config.Read()
			state := state{cfg: &cfg}
			username := "G0SU19O2"
			cmd := command{Name: tt.command, Args: []string{username}}
			if err := handlerLogin(&state, cmd); err != nil {
				t.Errorf("Fail to handler login with %v", err)
			}
			if cfg.CurrentUserName != username {
				t.Errorf("Want username: %s, got %s instead", username, cfg.CurrentUserName)
			}
		})
	}
}
