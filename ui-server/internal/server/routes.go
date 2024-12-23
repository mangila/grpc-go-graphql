package server

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"ui-server/web"
)

// Simple Middleware for logging
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Request: %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // Call the next handler in the chain
	})
}
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
