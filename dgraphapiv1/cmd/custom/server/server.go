package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Server holds the custom server router and exposes the
// required helper methods.
type Server struct {
	httpServer *http.Server
}

// New generates a pointer instance of the Server object.
func New(middlware mux.MiddlewareFunc, usersHandler, specimensHandler http.HandlerFunc) *Server {
	router := mux.NewRouter()

	router.Use(middlware)

	router.HandleFunc("/users", usersHandler)
	router.HandleFunc("/specimens", specimensHandler)

	customServer := &http.Server{
		Addr:         "custom:4080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return &Server{
		httpServer: customServer,
	}
}

// Start starts the configured server.
func (s *Server) Start() {
	s.httpServer.ListenAndServe()
}

// Stop stops the configured server.
func (s *Server) Stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}
