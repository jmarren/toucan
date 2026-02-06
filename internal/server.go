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

	router := &Router{Mux: mux}
	// https://upload.wikimedia.org/wikipedia/commons/thumb/f/f0/Chess_kdt45.svg/60px-Chess_kdt45.svg.png
	router.Page("GET /", controllers.Home)

	// create server
	s := &http.Server{
		Addr:    ":6060",
		Handler: router.Mux,
	}

	log.Printf("listening on %s\n", s.Addr)
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Fprintf(os.Stderr, "error listening and serving: %s\n", err)
	}

}
