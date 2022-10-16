package controller

import (
	"encoding/json"
	"net/http"
	"sesi4/server/view"
)

func WriteJsonResponse(w http.ResponseWriter, payload *view.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(payload.Status)
	json.NewEncoder(w).Encode(payload)
}
