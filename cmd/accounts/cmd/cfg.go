package cmd

import (
	"os"
)

var cfg = New()

// DB ...
type DB struct {
	Username string
	Password string
	Port     int
	Database string
}

// Nats ...
type Nats struct {
	Credentials string
	Url         string
}

// Flags contains the command line flags.
type Flags struct {
	Addr string
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
