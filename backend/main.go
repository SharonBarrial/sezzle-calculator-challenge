package main

import (
	"log"
	"net/http"
	"os"

	"github.com/SharonBarrial/calculator-backend/handlers"
)

// corsMiddleware allows the React dev server (a different origin) to call
// this API. In a real production deployment this would be locked down to
// the actual frontend origin instead of "*|"
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /api/add", handlers.Add)
	mux.HandleFunc("POST /api/subtract", handlers.Subtract)
	mux.HandleFunc("POST /api/multiply", handlers.Multiply)
	mux.HandleFunc("POST /api/divide", handlers.Divide)
	mux.HandleFunc("POST /api/power", handlers.Power)
	mux.HandleFunc("POST /api/sqrt", handlers.Sqrt)
	mux.HandleFunc("POST /api/percentage", handlers.Percentage)
	mux.HandleFunc("GET /api/health", handlers.Health)

	return corsMiddleware(mux)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("calculator backend listening on :%s", port)
	if err := http.ListenAndServe(":"+port, routes()); err != nil {
		log.Fatal(err)
	}
}
