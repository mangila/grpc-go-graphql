package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"shared/logger"
	"ui-server/web"
)

// loggingMiddleware - Simple Middleware for logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Logger.Infof("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}

// RegisterRoutes - register server routes and add web handlers
func (s *Server) RegisterRoutes() http.Handler {
	r := mux.NewRouter()
	fileServer := http.FileServer(http.FS(web.Files))
	r.NotFoundHandler = http.HandlerFunc(web.NotFoundHandler)
	r.MethodNotAllowedHandler = http.HandlerFunc(web.MethodNotAllowedHandler)
	r.PathPrefix("/assets/").Handler(fileServer)
	r.HandleFunc("/", web.HomeHandler).Methods(http.MethodGet)
	r.HandleFunc("/order", web.OrderHandler).Methods(http.MethodGet)
	r.HandleFunc("/product", web.ProductHandler).Methods(http.MethodGet)
	r.HandleFunc("/customer", web.CustomerHandler).Methods(http.MethodGet)
	r.HandleFunc("/hello", web.HelloHandler).Methods(http.MethodPost)
	r.Use(loggingMiddleware)
	return r
}
