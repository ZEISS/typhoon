package config

import (
	"fmt"
	"os"
)

// Database ...
type Database struct {
	Addr     string
	Database string
	Password string
	Port     int
	Prefix   string
	Username string
}

// Flags contains the command line flags.
type Flags struct {
	Addr     string
	Database *Database
}

// DSN for PostgreSQL.
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.Flags.Database.Addr, c.Flags.Database.Username, c.Flags.Database.Password, c.Flags.Database.Database, c.Flags.Database.Port)
}

// NewFlags ...
func NewFlags() *Flags {
	return &Flags{
		Database: &Database{
			Addr:     "host.docker.internal",
			Database: "example",
			Password: "example",
			Port:     5432,
			Username: "example",
			Prefix:   "typhoon_",
		},
	}
}

// Prefix ...
func (c *Config) Prefix() string {
	return c.Flags.Database.Prefix
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
