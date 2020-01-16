package server

import (
	"github.com/Phisto/eventusfileserver/config"
	"github.com/Phisto/eventusfileserver/server/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
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
	s.setRouters()
}

func (s *Server) setMiddleware() {
	// tell the ruter which middleware to use
	s.Router.Use(
		// used to log the request to the console | development
		middleware.Logger,
		// helps to redirect wrong requests (why do one want that?)
		//middleware.RedirectSlashes,
		// tries to recover after panics (?)
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
func (s *Server) setRouters() {

	// GET requests
	s.Router.Get("/images/{imageIdentifier}", s.handleRequest(handler.Download))
	s.Router.Get("/status", s.handleRequest(handler.Status))
	s.Router.Get("/status/files", s.handleRequest(handler.Files))

	// POST requests
	s.Router.Post("/images/upload", s.handleRequest(handler.MultipartUpload))

	// PATCH
	s.Router.Patch("/images/{imageIdentifier}", s.handleRequest(handler.Update))

	// s.handleRequest(handler.UpdateFestival)
	//s.Router.Get("/festivals", s.handleRequest(handler.GetFestivals)) // all or ?ids=1;2;3
	//s.Router.Get("/artists", s.handleRequest(handler.GetArtists))

	/*
		// POST requests
		s.Router.Post("/festivals", s.handleRequest(handler.CreateFestival))
		s.Router.Post("/artists", s.handleRequest(handler.CreateArtist))

		s.Router.Post("/festivals/{objectID}/image/{resourceID}", s.handleRequest(handler.SetImageForFestival))
		s.Router.Post("/festivals/{objectID}/links/{resourceID}", s.handleRequest(handler.SetLinkForFestival))

		// PATCH requests
		s.Router.Patch("/festivals/{objectID}", s.handleRequest(handler.UpdateFestival))
		s.Router.Patch("/artists/{objectID}", s.handleRequest(handler.UpdateArtist))


		// DELETE requests
		s.Router.Delete("/festivals/{objectID}", s.handleRequest(handler.DeleteFestival))
		s.Router.Delete("/artists/{objectID}", s.handleRequest(handler.DeleteArtist))
	*/
}

// Run the server on it's router
func (s *Server) Run(host string) {
	log.Fatal(http.ListenAndServe(host, s.Router))
}

// function prototype to inject config instance in handleRequest()
type RequestHandlerFunction func(config *config.Config, w http.ResponseWriter, r *http.Request)

// inject DB in handler functions
func (s *Server) handleRequest(handler RequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(s.Config, w, r)
	}
}
