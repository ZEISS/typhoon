package cmd

import (
	"fmt"
	"os"
)

var cfg = New()

// DB ...
type DB struct {
	Addr     string
	Database string
	Password string
	Port     int
	Username string
}

// Flags contains the command line flags.
type Flags struct {
	Addr string
	DB   *DB
}

// DSN for PostgreSQL.
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.Flags.DB.Addr, c.Flags.DB.Username, c.Flags.DB.Password, c.Flags.DB.Database, c.Flags.DB.Port)
}

// NewFlags ...
func NewFlags() *Flags {
	return &Flags{
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
