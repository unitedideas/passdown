package handler

import (
	"encoding/json"
	"net/http"
	"passdown/server/src/db/model"
	"passdown/server/src/main/db"
)

func HandleGetShiftsRequest(database *db.Database, w http.ResponseWriter, r *http.Request) {
	// This is an example. You'll replace the hardcoded shifts with a DB query
	shifts := []model.Shift{} // assuming you have model.Shift structure already defined

	err := database.Connection.Select(&shifts, "SELECT * FROM shifts")
	if err != nil {
		http.Error(w, "Failed to get shifts", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shifts)
}
