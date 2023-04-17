package api

import (
	"encoding/json"
	"net/http"
)

func GetAPIKey(w http.ResponseWriter, r *http.Request, name string) {
	json.NewEncoder(w).Encode(name)
}
