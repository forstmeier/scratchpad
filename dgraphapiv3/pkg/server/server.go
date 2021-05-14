package server

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"github.com/specimenguru/api/pkg/database/blevedb"
	"github.com/specimenguru/api/pkg/handlers"
)

const databaseFile = "specimens.bleve"

// Server exposes the HTTP handlers for the root API.
type Server struct {
	httpServer *http.Server
}

// New generates a Server pointer instance.
func New() (*Server, error) {
	router := mux.NewRouter()

	db, err := blevedb.New(databaseFile)
	if err != nil {
		return nil, &ErrorNewDatabase{err: err}
	}

	client, err := handlers.NewClient(db)
	if err != nil {
		return nil, &ErrorNewHandlersClient{err: err}
	}

	router.HandleFunc("/database", client.AddHandler()).Methods("POST")

	router.HandleFunc("/database", client.QueryHandler()).Methods("GET")

	return &Server{
		httpServer: &http.Server{
			Addr:         "localhost:8000",
			Handler:      router,
			ReadTimeout:  10 * time.Second,
			WriteTimeout: 10 * time.Second,
		},
	}, nil
}

// Start starts the configured server.
func (s *Server) Start() {
	s.httpServer.ListenAndServe()
}

// Stop stops the configured server.
func (s *Server) Stop(ctx context.Context) {
	s.httpServer.Shutdown(ctx)
}
