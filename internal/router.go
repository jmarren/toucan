package internal

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/jmarren/toucan/internal/views"
)

type Router struct {
	Mux *http.ServeMux
}

// apply middlewares
func (r *Router) Handle(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	handler = SessionMiddleware(handler)
	r.Mux.Handle(path, http.HandlerFunc(handler))
}

// handles a page component
func (r *Router) Page(path string, ch func(w http.ResponseWriter, r *http.Request) templ.Component) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		component := ch(w, r)
		// if hx-request render only the component
		// otherwise render inside the root page component
		if r.Header.Get("HX-Request") == "true" {
			component.Render(r.Context(), w)
		} else {
			views.Page(component).Render(r.Context(), w)
		}
	}

	r.Handle(path, handler)
}

// handles a component with the given method, path, and componentHandler
func (r *Router) Component(method string, path string, ch func(w http.ResponseWriter, r *http.Request) templ.Component) {
	fullPath := method + " /components" + path

	handler := func(w http.ResponseWriter, r *http.Request) {
		ch(w, r).Render(r.Context(), w)
	}

	r.Handle(fullPath, handler)
}
