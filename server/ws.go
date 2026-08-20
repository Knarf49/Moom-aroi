package main

import (
	"encoding/json"
	"log"
	"server/controller"
	wsevents "server/controller/event"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/contrib/v3/websocket/event"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func registerWs(app *fiber.App, db *gorm.DB) {
	app.Use("/ws", func(c fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next() //execute method ต่อไปนี้ (event.On ข้างล่าง)
		}
		return fiber.ErrUpgradeRequired
	})

	event.On(event.EventMessage, func(ep *event.EventPayload) {
		var message struct {
			Type    string          `json:"type"`
			Payload json.RawMessage `json:"payload"`
		}
		if err := json.Unmarshal(ep.Data, &message); err != nil {
			return
		}

		ep.Data = message.Payload
		switch message.Type {
			case "create_order":
				wsevents.Create_order(db, ep)
			case "order_food":
				wsevents.Order_food(db, ep)
		}
	})

	app.Get("/ws/:id", controller.TableIdRequired(db), event.New(func(kws *event.Websocket) {
		log.Println("Client connected:", kws.GetUUID())
	}))
}
