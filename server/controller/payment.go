// TODO: add validate tableId
package controller

import (
	"encoding/base64"
	"fmt"
	"os"
	models "server/schema"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	godotenv "github.com/joho/godotenv"
	promptpayqr "github.com/kazekim/promptpay-qr-go"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func Checkout(c fiber.Ctx, db *gorm.DB) error {
	var req struct {
		OrderID uint      `json:"orderId"`
		TableID uuid.UUID `json:"tableId"`
	}

	err := godotenv.Load()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Error loading .env",
		})
	}
	phone := os.Getenv("PROMPTPAY_PHONE")
	/*
		1. connect order_items to the dishes with dish_id
		2. filter only order_id that equal to input orderId
		3. SUM(order_items.quantity * dishes.price)
		4. wrap COALESCE around SUM to handle the case that row is empty so it convert NULL -> 0
	*/
	// 1. Parse the incoming request
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid request body",
		})
	}

	// 2. Ensure OrderID is valid before querying
	if req.OrderID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"error":   "Invalid order ID",
		})
	}

	// 3. Calculate the total price using a SQL JOIN
	//ลำดับการเรียก method ไม่มีผล เรียก sum ก่อน join ได้
	var totalPrice float64
	err = db.Transaction(func(tx *gorm.DB) error {
		var order models.Order
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&order, req.OrderID).Error; err != nil {
			return err
		}
		if order.TableID != req.TableID {
			return fmt.Errorf("table mismatch")
		}
		if order.Status != "pending" {
			return fmt.Errorf("order not awaiting checkout")
		}
		if err := tx.Table("order_items").
			Select("COALESCE(SUM(order_items.quantity * dishes.price), 0)").
			Joins("JOIN dishes ON dishes.id = order_items.dish_id").
			Where("order_items.order_id = ?", req.OrderID).
			Scan(&totalPrice).Error; err != nil {
			return err
		}
		if totalPrice <= 0 {
			return fmt.Errorf("order total is zero")
		}
		return tx.Model(&order).Update("status", "awaiting_payment").Error
	})
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "error": err.Error()})
	}

	qr, err := promptpayqr.QRForTargetWithAmount(phone, fmt.Sprintf("%.2f", totalPrice))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "error": "Error generating PromptPay QR code"})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"total":   totalPrice,
		"qr_code": fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(*qr)),
	})
}

// func UploadSlip(c fiber.Ctx, db *gorm.DB) error{
// 	var req struct {
// 		SlipImg string
// 		OrderID uint    `json:"orderId"`
// 		TableID uuid.UUID    `json:"tableId"`
// 	}

// }
