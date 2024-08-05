package cmd

import (
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
	Prefix   string
}

// Flags contains the command line flags.
type Flags struct {
	Addr                string `envconfig:"TYPHOON_ACCOUNTS_ADDR" default:":8084"`
	DatabaseURI         string `envconfig:"TYPHOON_ACCOUNTS_DATABASE_URI" default:""`
	DatabaseTablePrefix string `envconfig:"TYPHOON_ACCOUNTS_DATABASE_TABLE_PREFIX" default:"typhoon_"`
}

// NewFlags ...
func NewFlags() *Flags {
	return &Flags{}
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
