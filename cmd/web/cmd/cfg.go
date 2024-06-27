package cmd

import (
	"fmt"
	"os"
)

var cfg = New()

// DB ...
type DB struct {
	Addr     string `envconfig:"TYPHOON_DB_ADDR" default:"host.docker.internal"`
	Username string `envconfig:"TYPHOON_DB_USERNAME" default:"root"`
	Password string `envconfig:"TYPHOON_DB_PASSWORD" default:""`
	Database string `envconfig:"TYPHOON_DB_DATABASE" default:"defaultdb"`
	Port     int    `envconfig:"TYPHOON_DB_PORT" default:"26257"`
}

// Flags contains the command line flags.
type Flags struct {
	Addr string `envconfig:"TYPHOON_ADDR" default:":3000"`
	Nats *Nats
	FGA  *FGA
	DB   *DB
}

// FGA contains the OpenFGA configuration.
type FGA struct {
	// ApiUrl ...
	ApiUrl string `envconfig:"TYPHOON_FGA_API_URL" default:"http://host.docker.internal:8080"`
	// StoreId ...
	StoreID string `envconfig:"TYPHOON_FGA_STORE_ID" default:""`
	// AuthorizationModelId ...
	AuthorizationModelID string `envconfig:"TYPHOON_FGA_AUTHORIZATION_MODEL_ID" default:""`
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
		FGA:  &FGA{},
		Nats: &Nats{},
		DB: &DB{
			Addr:     "host.docker.internal",
			Database: "defaultdb",
			Password: "",
			Port:     26257,
			Username: "root",
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
