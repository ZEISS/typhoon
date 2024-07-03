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
	Addr                    string `envconfig:"TYPHOON_WEB_ADDR" default:":3000"`
	DatabaseURI             string `envconfig:"TYPHOON_WEB_DATABASE_URI" default:""`
	DatabaseTablePrefix     string `envconfig:"TYPHOON_WEB_DATABASE_TABLE_PREFIX" default:"typhoon_"`
	FGAApiUrl               string `envconfig:"TYPHOON_WEB_FGA_API_URL" default:"http://host.docker.internal:8080"`
	FGAStoreID              string `envconfig:"TYPHOON_WEB_FGA_STORE_ID" default:""`
	FGAAuthorizationModelID string `envconfig:"TYPHOON_WEB_FGA_AUTHORIZATION_MODEL_ID" default:""`
	GothGitbubKey           string `envconfig:"TYPHOON_WEB_GITHUB_KEY" default:""`
	GothGithubSecret        string `envconfig:"TYPHOON_WEB_GITHUB_SECRET" default:""`
	GothGithubCallback      string `envconfig:"TYPHOON_WEB_GITHUB_CALLBACK" default:"http://localhost:3000/auth/github/callback"`
}

// NewFlags ...
func NewFlags() *Flags {
	return &Flags{
		Addr:                ":3000",
		DatabaseURI:         "host=host.docker.internal user=example password=example dbname=example port=5432 sslmode=disable",
		DatabaseTablePrefix: "typhoon_",
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
