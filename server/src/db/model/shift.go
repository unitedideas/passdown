package model

type Shift struct {
	ID             int64  `json:"id"`
	UserID         int64  `json:"user_id"`
	PlantID        int64  `json:"plant_id"`
	MachineDetails string `json:"machine_details"`
	Notes          string `json:"notes"`
	ShiftTime      string `json:"shift_time"`
}
