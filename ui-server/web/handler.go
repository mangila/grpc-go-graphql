package web

import (
	"github.com/a-h/templ"
	"log"
	"net/http"
	"ui-server/web/pages"
	"ui-server/web/partials"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	templ.Handler(pages.Error("404, this is not the page you are looking for")).ServeHTTP(w, r)
}

func MethodNotAllowedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	templ.Handler(pages.Error("method is not supported")).ServeHTTP(w, r)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(pages.Home()).ServeHTTP(w, r)
}

func OrderHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(pages.Order()).ServeHTTP(w, r)
}

func ProductHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(pages.Product()).ServeHTTP(w, r)
}

func CustomerHandler(w http.ResponseWriter, r *http.Request) {
	templ.Handler(pages.Customer()).ServeHTTP(w, r)
}

// HelloHandler - handles /hello post
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
	name := r.FormValue("name")
	component := partials.HelloPost(name)
	err = component.Render(r.Context(), w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Fatalf("Error rendering in HelloWebHandler: %e", err)
	}
}
