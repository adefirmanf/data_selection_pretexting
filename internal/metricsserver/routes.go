package metricsserver

import (
	"net/http"

	promhttp "github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-chi/chi"
)

// App .
type App struct{}

func (a *App) metricsRouter() http.Handler {
	r := chi.NewRouter()
	r.HandleFunc("/metrics", promhttp.Handler().(http.HandlerFunc))
	return r
}

// NewRouterMetrics .
func NewRouterMetrics() *App {
	return &App{}
}
