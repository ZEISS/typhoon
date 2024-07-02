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
	Prefix   string `envconfig:"TYPHOON_DB_PREFIX" default:"typhoon_"`
}

// Flags contains the command line flags.
type Flags struct {
	Addr string `envconfig:"TYPHOON_ADDR" default:":3000"`
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

// DSN for PostgreSQL.
func (c *Config) DSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", c.Flags.DB.Addr, c.Flags.DB.Username, c.Flags.DB.Password, c.Flags.DB.Database, c.Flags.DB.Port)
}

// NewFlags ...
func NewFlags() *Flags {
	return &Flags{
		FGA: &FGA{},
		DB: &DB{
			Addr:     "host.docker.internal",
			Database: "defaultdb",
			Password: "",
			Port:     26257,
			Username: "root",
			Prefix:   "typhoon_",
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

// Prefix ...
func (c *Config) Prefix() string {
	return c.Flags.DB.Prefix
}

// Cwd returns the current working directory.
func (c *Config) Cwd() (string, error) {
	return os.Getwd()
}
