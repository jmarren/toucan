package internal

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jmarren/toucan/internal/views"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	views.Page().Render(r.Context(), w)
	// w.Write([]byte("hi!"))
}

func Start() {
	mux := http.NewServeMux()

	fs := http.FileServer(http.Dir("./public"))

	mux.Handle("GET /public", fs)

	mux.Handle("GET /", http.HandlerFunc(rootHandler))

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
