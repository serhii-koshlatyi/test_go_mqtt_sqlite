package config

import "github.com/BurntSushi/toml"

// Mosquitto represent config for Mosquitto mqtt broker
type Mosquitto struct {
	Broker   string `toml:"broker"`
	Port     int    `toml:"port"`
	Topic    string `toml:"topic"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

// SQLite represent config for SQLite DB
type SQLite struct {
	DSN string `toml:"dsn"`
}

// Config represents mix of settings for the app.
type Config struct {
	Mosquitto Mosquitto `toml:"mqtt"`
	SQLite    SQLite    `toml:"sqlite"`
}

// ReadConfig creates reads application configuration from the file.
func ReadConfig(path string) (*Config, error) {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
