package service

import (
	"net/http"

	repo "github.com/anhdvu/blueprint/internal/repository"
	"github.com/anhdvu/blueprint/pkg/aulog"
	"github.com/go-chi/chi/v5"
)

// A Service represents a web service or a web app
type Service struct {
	config config
	data   *repo.Repositories
	log    *aulog.Logger
	router *chi.Mux
}

type config struct {
	env    string
	port   int
	secret string
}

// Implement http.Handler interface for service
func (s *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
