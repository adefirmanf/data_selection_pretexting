package httpserver

import (
	"net/http"

	"github.com/adefirmanf/go-api-template/internal/httpserver/handlers"

	"github.com/go-chi/chi"
)

// App .
type App struct {}

func (a *App) init() http.Handler {
	r := chi.NewRouter()

	echoMessage := handlers.NewEchoMessageHandler()

	r.Route("/{message}", func(r chi.Router) {
		r.Use(handlers.ContextMessage)
		r.Get("/", echoMessage.GetHelloMessage)
	})

	return r
}

// NewRouterApp .
func NewRouterApp() *App {
	return &App{}
}
