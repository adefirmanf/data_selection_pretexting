package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
)

const (
	contextMessageKey = contextKey("message")
)

// ContextMessage .
func ContextMessage(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		params := chi.URLParam(r, "message")
		ctx := context.WithValue(r.Context(), contextMessageKey, params)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

type contextKey string

func (c contextKey) String() string {
	return string(c)
}
