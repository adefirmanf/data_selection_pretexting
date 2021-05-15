package config

import "fmt"

// Config .
type Config struct {
	c IConfig
}

var defaultConfig = new(Config)

// ErrNoInitializationConfig .
var ErrNoInitializationConfig = fmt.Errorf("config: %s", "No driver is loaded, use blank import to load driver to config.")

// Environment .
func (c *Config) Environment() string {
	return c.c.Environment()
}

// AppPort .
func (c *Config) AppPort() int {
	return c.c.AppPort()
}

// MetricsPort .
func (c *Config) MetricsPort() int {
	return c.c.MetricsPort()
}

// PGConfigConnectionString .
func (c *Config) PGConfigConnectionString() string {
	return c.c.PGConfigConnectionString()
}

// BearerTokenTwitter .
func (c *Config) BearerTokenTwitter() string {
	return c.c.BearerTokenTwitter()
}

// Init .
func Init(c IConfig) {
	defaultConfig.c = c
}

// Load .
func Load() *Config {
	if defaultConfig.c == nil {
		panic(ErrNoInitializationConfig)
	}
	return defaultConfig
}
