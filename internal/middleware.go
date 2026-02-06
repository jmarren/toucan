package internal

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/jmarren/toucan/internal/models"
)

func createSessionId() string {
	return uuid.New().String()
}

// check for session cookie and add it if not present
func SessionMiddleware(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		var sessionId = ""
		sessionCookie, err := r.Cookie("session")

		if err != nil {
			sessionId = createSessionId()
			cookie := &http.Cookie{
				Name:  "session",
				Value: sessionId,
			}
			http.SetCookie(w, cookie)
		} else {
			sessionId = sessionCookie.Value
		}
		r = r.WithContext(context.WithValue(r.Context(), "sessionId", sessionId))

		board, err := models.GetBoard(sessionId)

		rowZeroColTwo := board.GetSquare(0, 2)
		fmt.Printf("[0, 2]: %s\n", rowZeroColTwo)

		if err != nil {
			fmt.Printf("error: %s\n", err)
		} else {
			fmt.Printf("board: %v\n", board)
		}

		handler(w, r)
	}
}
