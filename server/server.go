package server

import (
	"log"
	"net/http"

	"github.com/Festivals-App/festivals-fileserver/server/config"
	"github.com/Festivals-App/festivals-fileserver/server/handler"
	"github.com/Festivals-App/festivals-identity-server/authentication"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// Server has router and db instances
type Server struct {
	Router *chi.Mux
	Config *config.Config
}

// Initialize the server with predefined configuration
func (s *Server) Initialize(config *config.Config) {

	// create router
	s.Router = chi.NewRouter()

	// set config
	s.Config = config

	// prepare  Router
	s.setMiddleware()
	s.setWalker()
	s.setRoutes()
}

func (s *Server) setMiddleware() {
	// tell the ruter which middleware to use
	s.Router.Use(
		// used to log the request to the console | development
		middleware.Logger,
		// tries to recover after panics
		middleware.Recoverer,
	)
}

func (s *Server) setWalker() {

	walkFunc := func(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
		log.Printf("%s %s \n", method, route)
		return nil
	}
	if err := chi.Walk(s.Router, walkFunc); err != nil {
		log.Panicf("Logging err: %s\n", err.Error())
	}
}

// setRouters sets the all required routers
func (s *Server) setRoutes() {

	s.Router.Get("/health", s.handleRequestWithoutAuthentication(handler.GetHealth))
	s.Router.Get("/version", s.handleRequestWithoutAuthentication(handler.GetVersion))
	s.Router.Get("/info", s.handleRequestWithoutAuthentication(handler.GetInfo))
	s.Router.Get("/status", s.handleRequestWithoutAuthentication(handler.Status))
	s.Router.Get("/files", s.handleRequestWithoutAuthentication(handler.Files))

	// GET requests
	s.Router.Get("/images/{imageIdentifier}", s.handleRequest(handler.Download))
	s.Router.Get("/pdf/{pdfIdentifier}", s.handleRequest(handler.DownloadPDF))

	// POST requests
	s.Router.Post("/images/upload", s.handleRequest(handler.MultipartUpload))
	s.Router.Post("/pdf/upload", s.handleRequest(handler.MultipartPDFUpload))

	// PATCH
	s.Router.Patch("/images/{imageIdentifier}", s.handleRequest(handler.Update))
	s.Router.Patch("/pdf/{pdfIdentifier}", s.handleRequest(handler.UpdatePDF))
}

func (s *Server) setFestivaslFilesAPIRoutes() {

}

// Run the server on it's router
func (s *Server) Run(host string) {
	//log.Fatal(http.ListenAndServeTLS(host, "/cert", "/keys", s.Router))
	log.Fatal(http.ListenAndServe(host, s.Router))
}

// function prototype to inject config instance in handleRequest()
type RequestHandlerFunction func(config *config.Config, w http.ResponseWriter, r *http.Request)

// inject Config in handler functions
func (s *Server) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {

	return authentication.IsAuthenticated(s.Config.APIKeys, func(w http.ResponseWriter, r *http.Request) {
		handler(s.Config, w, r)
	})
}

// inject Config in handler functions
func (s *Server) handleRequestWithoutAuthentication(requestHandler RequestHandlerFunction) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestHandler(s.Config, w, r)
	})
}
