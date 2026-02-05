package internal

import "net/http"

type Router struct {
	Mux *http.ServeMux
}

func (r *Router) Handle(path string, handler func(w http.ResponseWriter, r *http.Request)) {
	handler = SessionMiddleware(handler)
	r.Mux.Handle(path, http.HandlerFunc(handler))
}
