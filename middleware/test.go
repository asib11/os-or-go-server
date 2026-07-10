package middleware

import (
	"log"
	"net/http"
)

func Test(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Test middleware logic
		next.ServeHTTP(w, r)
		log.Println("this is test middleware: test print hobo")
	})
}
