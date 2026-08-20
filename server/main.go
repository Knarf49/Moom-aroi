package main

import (
	"log"
	controller "server/controller"
	models "server/schema"
	utils "server/utils"
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/static"
)

// TODO: finish cart page (cart using local storage)
// TODO: implement kitchen page
// TODO: test checkout route
// TODO: implement slip upload route
func main() {
	db, err := utils.SetupDb()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New())

	menu := app.Group("/api/menu")
	menu.Get("/", func(c fiber.Ctx) error {
		dishId := c.Query("id")
		var dishes []models.Dish
		var err error
		if dishId == "" {
			dishes, err = controller.GetAllMenu(db)
		} else {
			id, parseErr := strconv.ParseUint(dishId, 10, 64)
			if parseErr != nil {
				return c.Status(400).JSON(fiber.Map{"error": "invalid id"})
			}
			dishes, err = controller.GetMenuById(db, uint(id))
		}
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(dishes)
	})
	menu.Get("/recommend", func(c fiber.Ctx) error {
		dishes, err := controller.GetRecommendMenu(db)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(dishes)
	})

	app.Post("/api/checkout", controller.TableIdRequired(db), func(c fiber.Ctx) error {
		return controller.Checkout(c, db)
	})

	// app.Post("/api/payment/slip", controller.TableIdRequired(db), func(c fiber.Ctx) error{
	// 	return controller.UploadSlip()
	// })

	registerWs(app, db)

	app.Get("/uploads/*", static.New("./uploads"))
	app.Get("/*", static.New("../client/build"))

	log.Println("Server running on http://localhost:8081")
	log.Fatal(app.Listen(":8081"))
}
