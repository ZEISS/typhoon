package cmd

import (
	"fmt"
	"os"
)

// DB ...
type DB struct {
	Addr     string `envconfig:"TYPHOON_DB_ADDR" default:"host.docker.internal"`
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

// DSN for PostgreSQL.
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.Flags.DB.Addr, c.Flags.DB.Username, c.Flags.DB.Password, c.Flags.DB.Database, c.Flags.DB.Port)
}

// Cwd returns the current working directory.
func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}
