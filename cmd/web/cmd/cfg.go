package cmd

import (
	"fmt"
	"os"
)

var cfg = New()

// DB ...
type DB struct {
	Addr     string `envconfig:"TYPHOON_DB_ADDR" default:"host.docker.internal"`
	Database string `envconfig:"TYPHOON_DB_DATABASE" default:"example"`
	Password string `envconfig:"TYPHOON_DB_PASSWORD" default:"example"`
	Port     int    `envconfig:"TYPHOON_DB_PORT" default:"5432"`
	Username string `envconfig:"TYPHOON_DB_USERNAME" default:"example"`
}

// Flags contains the command line flags.
type Flags struct {
	Addr string
	Nats *Nats
	DB   *DB
}

// Nats contains the NATS configuration.
type Nats struct {
	Credentials string `envconfig:"TYPHOON_NATS_CREDENTIALS" default:"sys.creds"`
	URL         string `envconfig:"TYPHOON_NATS_URL" default:"nats://localhost:4222"`
}

// DSN for PostgreSQL.
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.Flags.DB.Addr, c.Flags.DB.Username, c.Flags.DB.Password, c.Flags.DB.Database, c.Flags.DB.Port)
}

// NewFlags ...
func NewFlags() *Flags {
	return &Flags{
		Nats: &Nats{},
		DB: &DB{
			Addr:     "host.docker.internal",
			Database: "example",
			Password: "example",
			Port:     5432,
			Username: "example",
		},
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
