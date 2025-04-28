package config

import (
	"encoding/json"
	"os"
)
const configFileName = ".gatorconfig.json"
type Config struct {
	Db_url string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() (Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	configFile, err := os.Open(configPath)
	if err != nil {
		return Config{}, err
	}
	defer configFile.Close()
	var config Config
	if err = json.NewDecoder(configFile).Decode(&config); err != nil {
		return Config{}, err
	}
	return config, nil
}

func (c *Config) SetUser(username string) error {
	c.CurrentUserName = username
	return write(*c)
}
func write(cfg Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	file, err := os.Create(configPath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = json.NewEncoder(file).Encode(cfg)
	if err != nil {
		return err
	}
	return err
}
func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homePath + "/" + configFileName, nil
}