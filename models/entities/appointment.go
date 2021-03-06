package entities

import (
	"database/sql"
	"time"

	"github.com/jinzhu/gorm"
)

// Appointment represent an entity in database
type Appointment struct {
	gorm.Model

	Name  string
	Date  time.Time
	Hour  sql.NullTime
	Local sql.NullString
}
