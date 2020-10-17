package mappers

import (
	"github.com/mirusky-dev/simple-agenda/models/dtos"
	"github.com/mirusky-dev/simple-agenda/models/entities"
)

// AppointmentEntityToDto convert a entity to dto
func AppointmentEntityToDto(ent *entities.Appointment) (*dtos.Appointment, error) {
	dto := new(dtos.Appointment)

	if ent.Local.Valid {
		local := ent.Local.String
		dto.Local = &local
	}

	if ent.Hour.Valid {
		hour := ent.Hour.Time.Format("15:06")
		dto.Hour = &hour
	}

	dto.Date = ent.Date.Format("02/01/2006")
	dto.Name = ent.Name

	dto.ID = ent.ID

	return dto, nil
}
