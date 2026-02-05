package internal

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

func createSessionId() string {
	return uuid.New().String()
}

// check for X-Session-Id header and add it if not present
func SessionMiddleware(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		sessionId := r.Header.Get("X-Session-Id")
		if sessionId == "" {
			sessionId = createSessionId()
			r = r.WithContext(context.WithValue(r.Context(), "sessionId", sessionId))
			w.Header().Add("HX-Trigger", "{\"sessionId\": { \"id\": \""+sessionId+"\"}}")
			w.Header().Add("X-Session-Id", sessionId)
		}
		handler(w, r)
	}
}
