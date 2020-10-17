package dtos

// Appointment is a simple data-transfer-object
type Appointment struct {
	ID    uint    `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Date  string  `json:"date,omitempty"`
	Hour  *string `json:"hour,omitempty"`
	Local *string `json:"local,omitempty"`
}
