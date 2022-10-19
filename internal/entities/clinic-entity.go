package entities

type ClinicEntity struct {
	ID          int64   `json:"id"`
	ClinicName  string  `json:"clinic_name"`
	Address     string  `json:"address"`
	PhoneNumber string  `json:"phone_number"`
	Latitude    float64 `json:"latitude"`
	Longitude   float64 `json:"longitude"`
	Distance    float64 `json:"distance"`
}
