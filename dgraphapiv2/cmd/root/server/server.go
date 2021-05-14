package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Server provides the HTTP resources for the GraphQL "pass through"
// wrapper on the Dgraph API.
type Server struct {
	httpServer *http.Server
}

// New generates a Server pointer object instance.
func New(graphqlHandler http.HandlerFunc) *Server {
	router := mux.NewRouter()

	router.HandleFunc("/graphql", graphqlHandler)

	return &Server{
		httpServer: &http.Server{
			Addr:         "root:4080",
			Handler:      router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
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
