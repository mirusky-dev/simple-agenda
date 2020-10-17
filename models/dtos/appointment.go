package dtos

// Appointment ...
type Appointment struct {
	ID    uint
	Name  string
	Date  string
	Hour  *string
	Local *string
}
