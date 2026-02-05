package controllers

import (
	"net/http"

	"github.com/jmarren/toucan/internal/views"
)

var components = make(map[string]func(w http.ResponseWriter, r *http.Request))

func AddComponents(mux *http.ServeMux) {
	makeComponents()
	for path, comp := range components {
		mux.Handle(path, http.HandlerFunc(comp))
	}
}

func makeComponent(method string, path string, fn func(w http.ResponseWriter, r *http.Request)) {
	components[method+" /components"+path] = fn
}

func makeComponents() {
	makeComponent("GET", "/side-panel", SidePanelHandler)

}

func SidePanelHandler(w http.ResponseWriter, r *http.Request) {
	views.SidePanel().Render(r.Context(), w)
}

// type components struct {
// 	controllers map[string]func(w http.ResponseWriter, r *http.Request)
// }

// var Components = &components{
// 	controllers: make(map[string]func(w http.ResponseWriter, r *http.Request)),
// }

//	func init() {
//		Components.Handle("side-panel", SidePanelHandler)
//	}
//
//	func (c *components) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//		componentName := r.PathValue("component")
//		controller := c.controllers[componentName]
//		if controller == nil {
//			w.WriteHeader(http.StatusNotFound)
//			return
//		}
//		c.controllers[componentName](w, r)
//	}
// func (c *components) Handle(path string, controller func(w http.ResponseWriter, r *http.Request)) {
// 	c.controllers[path] = controller
// }
