package entities

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

// Appointment ...
type Appointment struct {
	gorm.Model

	Name  string
	Date  time.Time
	Hour  sql.NullTime
	Local sql.NullString
}
