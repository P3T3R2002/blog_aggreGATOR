package config

import (
	"encoding/json"
	"os"
	"log"
	"path/filepath"
)

const config_file_name = ".gatorconfig.json"

type Config struct {
	DB_URL           string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func CreateJson(dbURL string) (Config, error) {
	cfg := Config{}
	cfg.DB_URL = dbURL
	err := write(cfg)
	return cfg, err
}

func CheckJson() bool {
    fullPath := get_config_file_path()
	_, err := os.Stat(fullPath)
    return err == nil
}

func (cfg *Config)SetUser(user_name string) error {
	cfg.Current_user_name = user_name
	return write(*cfg)
}

func write(cfg Config) error {
	full_path := get_config_file_path()

	file, err := os.Create(full_path)
	if err != nil {
		return err
	}
	defer file.Close()
	
	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	full_path := get_config_file_path()

	file, err := os.Open(full_path)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	cfg := Config{}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil
}

func get_config_file_path() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Error geting HomeDir!")
	}
	full_path := filepath.Join(home, config_file_name)
	return full_path
}
