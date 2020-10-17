package dtos

// Appointment is a simple data-transfer-object
type Appointment struct {
	ID    uint
	Name  string
	Date  string
	Hour  *string
	Local *string
}
