package config

// IConfig we want to make sure the driver of env should implement IConfig to avoid
// missing implementation
type IConfig interface {
	Environment() string
	AppPort() int
	MetricsPort() int
	PGConfigConnectionString() string
	BearerTokenTwitter() string
}
