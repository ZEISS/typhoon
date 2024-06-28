package config

import (
	"fmt"
	"os"
)

var Cfg = New()

// DB ...
type DB struct {
	Addr     string `envconfig:"TYPHOON_DB_ADDR" default:"host.docker.internal"`
	Username string `envconfig:"TYPHOON_DB_USERNAME" default:"root"`
	Password string `envconfig:"TYPHOON_DB_PASSWORD" default:""`
	Database string `envconfig:"TYPHOON_DB_DATABASE" default:"defaultdb"`
	Port     int    `envconfig:"TYPHOON_DB_PORT" default:"26257"`
	SslMode  string `envconfig:"TYPHOON_DB_SSL_MODE" default:"disable"`
}

// Flags contains the command line flags.
type Flags struct {
	Addr string `envconfig:"TYPHOON_ADDR" default:":8084"`
	DB   *DB
}

// DSN for PostgreSQL.
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s connect_timeout=5", c.Flags.DB.Addr, c.Flags.DB.Username, c.Flags.DB.Password, c.Flags.DB.Database, c.Flags.DB.Port, c.Flags.DB.SslMode)
}

// NewFlags ...
func NewFlags() *Flags {
	return &Flags{
		Addr: ":8084",
		DB: &DB{
			Addr:     "host.docker.internal",
			Database: "defaultdb",
			Password: "",
			Port:     26257,
			Username: "root",
			SslMode:  "disable",
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
