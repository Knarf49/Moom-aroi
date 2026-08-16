package wsevents

import (
	"encoding/json"
	"log"
	models "server/schema"

	"github.com/gofiber/contrib/v3/websocket/event"
	"gorm.io/gorm"
)

func emit(ep *event.EventPayload, msg []byte) {
	if ep != nil && ep.Kws != nil && ep.Kws.Conn != nil {
		ep.Kws.Emit(msg, event.TextMessage)
	}
}

type createOrderPayload struct {
	TableID uint `json:"tableId"`
}

// payload : {MenuId: quantity, MenuId: quantity,.... }
func Create_order(db *gorm.DB, ep *event.EventPayload) {
	var payload createOrderPayload
	if err := json.Unmarshal(ep.Data, &payload); err != nil {
		emit(ep, []byte(`{"error":"invalid create_order payload"}`))
		return
	}
	newOrder := models.Order{TableID: payload.TableID}
	if err := db.Create(&newOrder).Error; err != nil {
		log.Printf("create order failed for table=%d: %v", payload.TableID, err)
		errMsg, _ := json.Marshal(map[string]any{
			"event": "create_order",
			"error": "failed to create order",
		})
		emit(ep, errMsg)
		return
	}

	ack, _ := json.Marshal(map[string]any{
		"event":   "create_order",
		"orderId": newOrder.ID,
	})
	emit(ep, ack)
}

type OrderFoodPayload struct {
	OrderMap map[uint]int `json:"orderMap"`
	TableID  uint         `json:"tableId"`
	OrderID  uint         `json:"orderId"`
}

// TODO: add error handler to functions below
// tx = db ที่เก็บชั่วคราวตอน query ก่อนที่จะเก็บถาวรตอนทุกอย่างสำเร็จ ไม่มี error
func Order_food(db *gorm.DB, ep *event.EventPayload) {
	var payload OrderFoodPayload
	// convert json into Go struct
	if err := json.Unmarshal(ep.Data, &payload); err != nil {
		emit(ep, []byte(`{"error":"invalid order_food payload"}`))
		return
	}
	if len(payload.OrderMap) == 0 {
		emit(ep, []byte(`{"error":"orderMap cannot be empty"}`))
		return
	}
	// log.Printf("order_food from %s: table=%s dishes=%v", ep.SocketUUID, payload.TableID, payload.MenuID)
	// loop through menuId and create OrderItem
	err := db.Transaction(func(tx *gorm.DB) error {
		// return อันนี้เป็นแค่ของ func(tx *gorm.DB) ไม่ใช่ของ func Order_food
		for id, quantity := range payload.OrderMap {
			newOrderItem := models.Order_items{Quantity: quantity, OrderID: payload.OrderID, DishID: id}
			if err := tx.Create(&newOrderItem).Error; err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("Order food failed for order=%d: %v", payload.OrderID, err)
		emit(ep, []byte(`{"event":"order_food","error":"failed to create order item"}`))
		return
	}

	// uint
	ack, _ := json.Marshal(map[string]any{
		"event":   "order_food",
		"status":  "received",
		"orderId": payload.OrderID,
	})
	emit(ep, ack)
}

type DeleteOrderPayload struct {
	Order models.Order `json:"order"`
}

func Delete_order(db *gorm.DB, ep *event.EventPayload) {
	var payload DeleteOrderPayload

	if err := json.Unmarshal(ep.Data, &payload); err != nil {
		emit(ep, []byte(`{"error":"invalid delete_order payload"}`))
		return
	}

	if err := db.Unscoped().Delete(&payload.Order).Error; err != nil {
		ack, _ := json.Marshal(map[string]any{
			"event":  "delete_order",
			"status": "failed to delete order",
		})
		emit(ep, ack)
		return
	}

	ack, _ := json.Marshal(map[string]any{
		"event":  "delete_order",
		"status": "deleted",
	})
	emit(ep, ack)
}
