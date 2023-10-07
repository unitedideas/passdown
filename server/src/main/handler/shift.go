package handler

import (
	"encoding/json"
	"net/http"
)

func HandleGetShiftsRequest(w http.ResponseWriter, r *http.Request) {
	shifts := []struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Shift string `json:"shift"`
	}{
		{ID: 1, Name: "John Doe", Shift: "Morning"},
		{ID: 2, Name: "Jane Smith", Shift: "Evening"},
		// Add more shifts as needed
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shifts)
}
