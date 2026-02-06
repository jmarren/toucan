package controllers

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/jmarren/toucan/internal/views"
)

func Home(w http.ResponseWriter, r *http.Request) templ.Component {
	return views.Home()
}
