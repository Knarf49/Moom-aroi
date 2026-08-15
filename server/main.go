package main

import (
	"log"
	controller "server/controller"
	gorm "server/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

func main() {
	db, err := gorm.SetupDb()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/api/menu", func(c *fiber.Ctx) error {
		dishes, err := controller.GetAllMenu(db)
		if err != nil {
			return c.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(dishes)
	})

	app.Use("/ws", func(c *fiber.Ctx) error {
		// เปิดใช้ websocket ถ้า request ที่ส่งมาถูกต้อง
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}

		return fiber.ErrUpgradeRequired
	})

	app.Get("/ws", websocket.New((func(c *websocket.Conn) {
		// websocket hub
		var (
			mt  int
			msg []byte
			err error
		)
		for {
			if mt, msg, err = c.ReadMessage(); err != nil {
				log.Println("Connection closed:", err)
			}
			log.Printf("Received: %s", msg)

			if err = c.WriteMessage(mt, msg); err != nil {
				log.Println("Write error:", err)
				break
			}
		}
	})))
	app.Static("/", "../client/build")

	log.Println("Server running on http://localhost:8081")
	log.Fatal(app.Listen(":8081"))
}
