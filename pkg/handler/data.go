package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"../repository"
)

// GetDataHandler handles the request for data.
func GetDataHandler(w http.ResponseWriter, r *http.Request) {
	// You can also get these from environment variables or config files
	const url = "https://your-api-url.com"
	const username = "yourUsername"
	const password = "yourPassword"

	data, err := repository.FetchData(url, username, password)
	if err != nil {
		log.Printf("Error fetching data: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Respond with JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
