package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// xMiddleware function that logs requests
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
		log.Printf("Completed in %v", time.Since(start))
	})
}

// Middleware to check API key in headers
func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		apiKey := r.Header.Get("X-API-KEY")
		if apiKey != "valid-api-key" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// Basic handler for home route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func intro(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "I'm Abner Mencia")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", loggingMiddleware(authMiddleware(http.HandlerFunc(homeHandler))))
	mux.Handle("/into", loggingMiddleware(http.HandlerFunc(intro)))

	fmt.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
