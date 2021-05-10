package metricsserver

import (
	"fmt"
	"net/http"
)

// Httpserver .
type Httpserver struct{}

// ListenAndServe .
func (H *Httpserver) ListenAndServe(port int) {
	r := NewRouterMetrics()
	http.ListenAndServe(fmt.Sprintf(":%v", port), r.metricsRouter())
}

// NewHTTPServer .
func NewHTTPServer() *Httpserver {
	return &Httpserver{}
}
