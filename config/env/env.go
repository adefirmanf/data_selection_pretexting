package env

import (
	"fmt"
	"os"
	"strconv"
)

// Env .
type Env struct{}

// Environment .
func (e *Env) Environment() string {
	return getStringOrDefault(EnvironmentEnv, "local")
}

// AppPort .
func (e *Env) AppPort() int {
	return getIntOrDefault(AppPort, 3000)
}

// MetricsPort .
func (e *Env) MetricsPort() int {
	return getIntOrDefault(MetricsPort, 9100)
}

// PGConfigConnectionString .
func (e *Env) PGConfigConnectionString() string {
	return getStringOrDefault(PGConfigConnectionString, "")
}

// BearerTokenTwitter .
func (e *Env) BearerTokenTwitter() string {
	return getStringOrDefault(BearerTokenTwitter, "")
}

// New .
func New() *Env {
	return &Env{}
}

func getStringOrDefault(key, def string) string {
	return getEnvOrDefault(key, def)
}

func getIntOrDefault(key string, def int) int {
	results := getEnvOrDefault(key, fmt.Sprint(def))
	i, err := strconv.ParseInt(results, 10, 64)
	if err != nil {
		return def
	}
	return int(i)
}

func getEnvOrDefault(key, def string) string {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	return v
}
