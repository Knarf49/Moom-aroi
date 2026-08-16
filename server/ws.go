package main

import (
	"log"
	wsevents "server/controller/event"

	"github.com/gofiber/contrib/v3/websocket"
	"github.com/gofiber/contrib/v3/websocket/event"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func registerWs(app *fiber.App, db *gorm.DB) {
	app.Use("/ws", func(c fiber.Ctx) error{
		if websocket.IsWebSocketUpgrade(c){
			c.Locals("allowed", true)
			return c.Next() //execute method ต่อไปนี้ (event.On ข้างล่าง)
		}
		return fiber.ErrUpgradeRequired
	})

	event.On("create_order", func(ep *event.EventPayload){
		wsevents.Create_order(db,ep)
	})

	event.On("order_food", func(ep *event.EventPayload) {
		wsevents.Order_food(db, ep)
	})
	event.On("delete_order", func(ep *event.EventPayload) {
		wsevents.Delete_order(db, ep)
	})

	app.Get("/ws/:id",event.New(func(kws *event.Websocket){
		log.Println("Client connected:",kws.GetUUID())
	}))
}