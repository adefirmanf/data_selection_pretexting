package httpserver

import (
	"fmt"
	"net/http"
)

// Httpserver .
type Httpserver struct {
	Metrics
}

// Metrics .
type Metrics struct {
}

// ListenAndServe .
func (h *Httpserver) ListenAndServe(port int) {
	r := NewRouterApp()
	http.ListenAndServe(fmt.Sprintf(":%v", port), r.init())
}

// NewHTTPServer .
func NewHTTPServer(h Httpserver) *Httpserver {
	return &Httpserver{
		Metrics: Metrics{},
	}
}
