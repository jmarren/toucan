package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmarren/toucan/internal/controllers"
)

func Start() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./public"))

	mux.Handle("GET /public/", http.StripPrefix("/public", fs))
	mux.Handle("GET /components/{component}", controllers.Components)
	mux.Handle("GET /", http.HandlerFunc(controllers.RootHandler))

	// create server
	s := &http.Server{
		Addr:    ":6060",
		Handler: mux,
	}

	log.Printf("listening on %s\n", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}

}
