package mappers

import (
	"database/sql"
	"errors"
	"time"

	"github.com/mirusky-dev/simple-agenda/models/dtos"
	"github.com/mirusky-dev/simple-agenda/models/entities"
)

// AppointmentDtoToEntity convert a dto to entity
func AppointmentDtoToEntity(dto *dtos.Appointment) (*entities.Appointment, error) {
	ent := new(entities.Appointment)
	if dto.Hour != nil {
		if hour, err := time.Parse("15:04", *dto.Hour); err == nil {
			ent.Hour = sql.NullTime{
				Time:  hour,
				Valid: true,
			}
		} else {
			return nil, errors.New("Invalid date-format for field hour, allowed HH:MM")
		}
	}

	if dto.Local != nil {
		ent.Local = sql.NullString{
			String: *dto.Local,
			Valid:  true,
		}
	}

	if dto.Name != "" {
		ent.Name = dto.Name
	}

	if date, err := time.Parse("02/01/2006", dto.Date); err == nil {
		ent.Date = date
	} else {
		return nil, errors.New("Invalid date-format for field hour, allowed DD/MM/YYYY")
	}

	return ent, nil
}

func NormalizeTime(dto *dtos.Appointment) {
	if dto.Hour != nil {
		hour, _ := time.Parse(time.RFC3339, *dto.Hour)
		hourStr := hour.Format("15:04")
		dto.Hour = &hourStr
	}

	if dto.Date != "" {
		date, _ := time.Parse(time.RFC3339, dto.Date)
		dto.Date = date.Format("02/01/2006")
	}

}
