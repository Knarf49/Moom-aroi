package main

import (
	"log"
	controller "server/controller"
	gorm "server/utils"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"
)

func main() {
	db, err := gorm.SetupDb()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/api/menu", func(c fiber.Ctx) error {
		dishes, err := controller.GetAllMenu(db)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(dishes)
	})
	
	registerWs(app, db)
	app.Get("/*", static.New("../client/build"))

	log.Println("Server running on http://localhost:8081")
	log.Fatal(app.Listen(":8081"))
}
