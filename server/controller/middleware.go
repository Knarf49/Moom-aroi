package controller

import (
	models "server/schema"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tableIdBody struct {
	TableId uuid.UUID `json:"tableId"`
}

func TableIdRequired(db *gorm.DB) fiber.Handler {
	return func(c fiber.Ctx) error {
		rawTableID := c.Params("id")
		var tableID uuid.UUID
		if rawTableID != "" {
			var err error
			tableID, err = uuid.Parse(rawTableID)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid or empty TableId!",
				})
			}
		} else {
			var body tableIdBody
			if err := c.Bind().Body(&body); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error": "Invalid request body!",
				})
			}
			tableID = body.TableId
		}

		if tableID == uuid.Nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or empty TableId!",
			})
		}

		var table models.Table
		if err := db.First(&table, tableID).Error; err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or empty TableId!",
			})
		}

		c.Locals("table", table)
		return c.Next()
	}
}
