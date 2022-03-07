package server

import (
	"net/http"

	"github.com/Festivals-App/festivals-fileserver/server/config"
	"github.com/Festivals-App/festivals-fileserver/server/handler"
	"github.com/Festivals-App/festivals-gateway/server/logger"
	"github.com/Festivals-App/festivals-identity-server/authentication"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog/log"
)

// Server has router and db instances
type Server struct {
	Router *chi.Mux
	Config *config.Config
}

// Initialize the server with predefined configuration
func (s *Server) Initialize(config *config.Config) {

	s.Router = chi.NewRouter()

	s.Config = config

	s.setMiddleware()
	s.setWalker()
	s.setRoutes()
}

func (s *Server) setMiddleware() {
	// tell the ruter which middleware to use
	s.Router.Use(
		// used to log the request to the console | development
		logger.Middleware(&log.Logger),
		// tries to recover after panics
		middleware.Recoverer,
	)
}

func (s *Server) setWalker() {

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Info().Msg(method + " " + route + " \n")
		return nil
	}
	if err := chi.Walk(s.Router, walkFunc); err != nil {
		log.Panic().Msg(err.Error())
	}
}

// setRouters sets the all required routers
func (s *Server) setRoutes() {

	s.Router.Get("/version", s.handleRequestWithoutAuthentication(handler.GetVersion))
	s.Router.Get("/info", s.handleRequestWithoutAuthentication(handler.GetInfo))
	s.Router.Get("/health", s.handleRequestWithoutAuthentication(handler.GetHealth))

	s.Router.Get("/log", s.handleAdminRequest(handler.GetLog))
	s.Router.Get("/status", s.handleAdminRequest(handler.Status))
	s.Router.Get("/files", s.handleAdminRequest(handler.Files))

	// GET requests
	s.Router.Get("/images/{imageIdentifier}", s.handleRequest(handler.Download))
	s.Router.Get("/pdf/{pdfIdentifier}", s.handleRequest(handler.DownloadPDF))

	// POST requests
	s.Router.Post("/images/upload", s.handleAdminRequest(handler.MultipartUpload))
	s.Router.Post("/pdf/upload", s.handleAdminRequest(handler.MultipartPDFUpload))

	// PATCH
	s.Router.Patch("/images/{imageIdentifier}", s.handleAdminRequest(handler.Update))
	s.Router.Patch("/pdf/{pdfIdentifier}", s.handleAdminRequest(handler.UpdatePDF))
}

// Run the server on it's router
func (s *Server) Run(host string) {
	if err := http.ListenAndServe(host, s.Router); err != nil {
		log.Fatal().Err(err).Msg("Startup failed")
	}
}

// function prototype to inject config instance in handleRequest()
type RequestHandlerFunction func(config *config.Config, w http.ResponseWriter, r *http.Request)

func (s *Server) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {

	return authentication.IsEntitled(s.Config.APIKeys, func(w http.ResponseWriter, r *http.Request) {
		handler(s.Config, w, r)
	})
}

func (s *Server) handleAdminRequest(requestHandler RequestHandlerFunction) http.HandlerFunc {

	return authentication.IsEntitled(s.Config.AdminKeys, func(w http.ResponseWriter, r *http.Request) {
		requestHandler(s.Config, w, r)
	})
}

// inject Config in handler functions
func (s *Server) handleRequestWithoutAuthentication(requestHandler RequestHandlerFunction) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestHandler(s.Config, w, r)
	})
}
