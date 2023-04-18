package api

import (
	"net/http"
	"encoding/json"
	"os"
)

func GetAPIKey(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(os.Getenv("GOOGLE_API_KEY"))
}
