package config

import (
	"os"
	"testing"
)

func TestGetConfigPath(t *testing.T) {
	path, err := getConfigFilePath()
	if err != nil {
		t.Fatalf("Failed to get config path: %v", err)
	}
	if path == "" {
		t.Error("Config path should not be empty")
	}
}

func TestConfigFileExist(t *testing.T) {
	configPath, err := getConfigFilePath()
	if err != nil {
		t.Fatalf("Failed to get config path: %v", err) 
	}
	if _, err := os.Stat(configPath); err != nil {
		t.Errorf("Config file should exist at %s: %v", configPath, err)
	}
}

func TestReadConfig(t *testing.T) {
	cfg, err := Read()
	if err != nil {
		t.Fatalf("Failed to read config: %v", err)
	}
	
	if cfg.Db_url == "" {
		t.Error("DB URL should not be empty")
	}
}

func TestSetAndGetUser(t *testing.T) {
	tests := []struct {
		name     string
		username string
	}{
		{"empty username", ""},
		{"valid username", "testuser"},
		{"special chars", "test@123"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cfg, err := Read()
			if err != nil {
				t.Fatalf("Failed to read config: %v", err)
			}

			if err := cfg.SetUser(tt.username); err != nil {
				t.Errorf("SetUser(%s) failed: %v", tt.username, err)
			}

			newCfg, err := Read()
			if err != nil {
				t.Fatalf("Failed to read updated config: %v", err)
			}

			if newCfg.CurrentUserName != tt.username {
				t.Errorf("Want username %s, got %s", tt.username, newCfg.CurrentUserName)
			}
		})
	}
}

func TestSetUser(t *testing.T) {
	cfg, err := Read()
	if err != nil {
		t.Error(err)
	}
	username := "G0SU19O2"
	err = cfg.SetUser("G0SU19O2")
	if err != nil {
		t.Errorf("Cannot perform set username, error: %q", err)
	}
	cfg, err = Read()
	if err != nil {
		t.Error(err)
	}
	if cfg.CurrentUserName != username {
		t.Errorf("Want username: %s, got %s instead", username, cfg.CurrentUserName)
	}
}
