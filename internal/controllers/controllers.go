package controllers

import (
	"net/http"

	"github.com/jmarren/toucan/internal/models"
	"github.com/jmarren/toucan/internal/views"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	users, _ := models.Query.GetAllUsers(r.Context())

	views.Page(users).Render(r.Context(), w)
}

// imports from views and models
