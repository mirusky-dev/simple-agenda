package handlers

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mirusky-dev/agenda-simple/database"
	"github.com/mirusky-dev/agenda-simple/models/dtos"
	"github.com/mirusky-dev/agenda-simple/models/entities"
	"github.com/mirusky-dev/agenda-simple/models/mappers"
	"gorm.io/gorm"
)

// AppointmentRouter setup routes for /appointments
func AppointmentRouter(app fiber.Router, gormDB *gorm.DB) {
	app.Post("/appointments", post(gormDB))
	app.Get("/appointments/:id", getByID(gormDB))
	app.Get("/appointments", getAll(gormDB))
	app.Put("/appointments/:id", put(gormDB))
	app.Delete("/appointments/:id", delete(gormDB))
}

// post handles Create operation
func post(gormDB *gorm.DB) fiber.Handler {
	db, _ := database.Clone(gormDB)

	return func(c *fiber.Ctx) error {
		var appointment = new(dtos.Appointment)
		err := c.BodyParser(&appointment)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "invalid request",
				"error":   err.Error(),
			})
		}

		ent, err := mappers.AppointmentDtoToEntity(appointment)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "invalid request",
				"error":   err.Error(),
			})
		}

		err = db.Create(ent).Error
		if err != nil {
			c.Status(fiber.StatusUnprocessableEntity)
			return c.JSON(fiber.Map{
				"message": "failed to persist data",
				"error":   err.Error(),
			})
		}

		dto, _ := mappers.AppointmentEntityToDto(ent)
		c.Response().Header.Set("Location", fmt.Sprintf(c.Path()+"/%+v", ent.ID))
		c.Status(fiber.StatusCreated)
		return c.JSON(fiber.Map{
			"message": "created",
			"error":   nil,
			"result":  dto,
		})
	}
}

// getByID handles Read operation by ID
func getByID(gormDB *gorm.DB) fiber.Handler {
	db, _ := database.Clone(gormDB)

	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "invalid request",
				"error":   err.Error(),
			})
		}

		var appointment = new(dtos.Appointment)
		err = db.Model(&entities.Appointment{}).Find(appointment, id).Error
		if err != nil || appointment.ID == 0 {
			if err == gorm.ErrRecordNotFound || appointment.ID == 0 {
				err = gorm.ErrRecordNotFound
				c.Status(fiber.StatusNotFound)
				return c.JSON(fiber.Map{
					"message": "not found data with the given id",
					"error":   err.Error(),
				})
			}
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "failed to find data",
				"error":   err.Error(),
			})
		}

		c.Status(fiber.StatusFound)
		return c.JSON(appointment)
	}
}

// getAll handles Read paginated operation
func getAll(gormDB *gorm.DB) fiber.Handler {
	db, _ := database.Clone(gormDB)

	return func(c *fiber.Ctx) error {
		var pager = new(dtos.Pager)
		c.QueryParser(pager)
		pager.Check()

		db = db.Limit(pager.Limit)
		db = db.Offset(pager.Offset)

		var appointment = new([]dtos.Appointment)
		err := db.Model(&entities.Appointment{}).Find(appointment).Error
		if err != nil {
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "failed to find data",
				"error":   err.Error(),
			})
		}

		return c.JSON(fiber.Map{
			"result": appointment,
			"_metadata": fiber.Map{
				"previous": pager.Previous(),
				"this":     pager.Page,
				"next":     pager.Next(),
				"limit":    pager.Limit,
				"offset":   pager.Offset,
			},
		})
	}
}

// put handles Update operation by ID
func put(gormDB *gorm.DB) fiber.Handler {
	db, _ := database.Clone(gormDB)

	return func(c *fiber.Ctx) error {

		idStr := c.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "invalid request",
				"error":   err.Error(),
			})
		}

		var appointmentUpdate = new(dtos.Appointment)
		err = c.BodyParser(&appointmentUpdate)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "invalid request",
				"error":   err.Error(),
			})
		}

		entUpdate, err := mappers.AppointmentDtoToEntity(appointmentUpdate)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "invalid request",
				"error":   err.Error(),
			})
		}
		entUpdate.ID = uint(id)
		err = db.Model(&entities.Appointment{}).Where("id = ?", id).Updates(entUpdate).Error
		if err != nil {
			c.Status(fiber.StatusUnprocessableEntity)
			return c.JSON(fiber.Map{
				"message": "failed to persist data",
				"error":   err.Error(),
			})
		}

		dto, _ := mappers.AppointmentEntityToDto(entUpdate)
		c.Response().Header.Set("Location", c.Path())
		c.Status(fiber.StatusCreated)
		return c.JSON(fiber.Map{
			"message": "updated",
			"error":   nil,
			"result":  dto,
		})
	}
}

// delete handles Delete operation by ID
func delete(gormDB *gorm.DB) fiber.Handler {
	db, _ := database.Clone(gormDB)

	return func(c *fiber.Ctx) error {

		idStr := c.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 64)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{
				"message": "invalid request",
				"error":   err.Error(),
			})
		}

		err = db.Delete(&entities.Appointment{}, id).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				c.Status(fiber.StatusNotFound)
				return c.JSON(fiber.Map{
					"message": "not found data with the given id",
					"error":   err.Error(),
				})
			}
			c.Status(fiber.StatusInternalServerError)
			return c.JSON(fiber.Map{
				"message": "failed to delete data",
				"error":   err.Error(),
			})
		}

		c.Status(fiber.StatusOK)
		return nil
	}
}
