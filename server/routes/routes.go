package routes

import (
	"net/http"

	"github.com/flexearl/halloween_2023.git/middleware"
)

func HandleRouting() {
	http.HandleFunc("/", withCORS(middleware.Home))
	http.HandleFunc("/register_user", withCORS(middleware.Register))
	http.HandleFunc("/get_user_puzzle_input", withCORS(middleware.GetUserPuzzleInput))
	http.HandleFunc("/check_user_pumpkins", withCORS(middleware.CheckUserPumpkins))
	http.HandleFunc("/add_pumpkin", withCORS(middleware.AddPumpkin))
	http.HandleFunc("/get_puzzle_content/", withCORS(middleware.GetPuzzleContent))
}

func withCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Allow frontend origin
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Handle preflight requests (OPTIONS)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	}
}
