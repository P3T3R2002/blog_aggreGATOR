package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const config_file_name = ".gatorconfig.json"

type Config struct {
	DB_URL           string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func (cfg *Config)Set_user(user_name string) error {
	cfg.Current_user_name = user_name
	return write(*cfg)
}

func write(cfg Config) error {
	full_path, err := get_config_file_path()
	if err != nil {
		return err
	}

	file, err := os.Open(full_path)
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
	full_path, err := get_config_file_path()
	if err != nil {
		return Config{}, err
	}

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

func get_config_file_path() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	full_path := filepath.Join(home, config_file_name)
	return full_path, nil
}
