package cmd

import (
	"os"
)

// DB ...
type DB struct {
	Username string `envconfig:"TYPHOON_DB_USERNAME" default:"example"`
	Password string `envconfig:"TYPHOON_DB_PASSWORD" default:"example"`
	Port     int    `envconfig:"TYPHOON_DB_PORT" default:"5432"`
	Database string `envconfig:"TYPHOON_DB_DATABASE" default:"example"`
}

// Nats ...
type Nats struct {
	Credentials string `envconfig:"TYPHOON_NATS_CREDENTIALS" default:"sys.creds"`
	Url         string `envconfig:"TYPHOON_NATS_URL" default:"nats://localhost:4222"`
}

// Flags contains the command line flags.
type Flags struct {
	DB   *DB
	Nats *Nats
}

// NewFlags ...
func NewFlags() *Flags {
	return &Flags{
		DB:   &DB{},
		Nats: &Nats{},
	}
}

// New ...
func New() *Config {
	return &Config{
		Flags: NewFlags(),
	}
}

// Config ...
type Config struct {
	Flags *Flags
}

// Cwd returns the current working directory.
func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}
