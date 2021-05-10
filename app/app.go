package app

// App holds all services required by application.
// i.e Logger, Metrics, or domain business
type App struct {
	*Metrics
}

// Metrics .
type Metrics struct {
}

// Logger .
type Logger struct{}

// MetricsApp .
func buildMetricsApp() *Metrics {
	return &Metrics{}
}

// New .
func New() *App {
	return &App{
		Metrics: buildMetricsApp(),
	}
}
