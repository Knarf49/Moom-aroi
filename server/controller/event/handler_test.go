// wsevents/handlers_test.go
package wsevents

import (
	"encoding/json"
	"testing"

	"github.com/glebarez/sqlite" // swap import if you're still on gorm.io/driver/sqlite
	"github.com/gofiber/contrib/v3/websocket/event"
	"gorm.io/gorm"

	models "server/schema"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open test db: %v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("failed to get sql.DB: %v", err)
	}
	sqlDB.SetMaxOpenConns(1) // ต้องมาก่อน Exec/Migrate ใดๆ ทั้งหมด

	if err := db.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
		t.Fatalf("failed to enable foreign keys: %v", err)
	}

	if err := db.AutoMigrate(&models.Table{}, &models.Dish{}, &models.Order{}, &models.Order_items{}); err != nil {
		t.Fatalf("failed to migrate: %v", err)
	}
	return db
}

func seedTable(t *testing.T, db *gorm.DB) models.Table {
	t.Helper()
	table := models.Table{Label: "Table 1"}
	if err := db.Create(&table).Error; err != nil {
		t.Fatalf("failed to seed table: %v", err)
	}
	return table
}

func seedDish(t *testing.T, db *gorm.DB, title string) models.Dish {
	t.Helper()
	dish := models.Dish{Title: title, Price: 50}
	if err := db.Create(&dish).Error; err != nil {
		t.Fatalf("failed to seed dish: %v", err)
	}
	return dish
}

func seedOrder(t *testing.T, db *gorm.DB, tableID uint) models.Order {
	t.Helper()
	order := models.Order{TableID: tableID}
	if err := db.Create(&order).Error; err != nil {
		t.Fatalf("failed to seed order: %v", err)
	}
	return order
}

func newTestPayload(t *testing.T, v any) *event.EventPayload {
	t.Helper()
	data, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("failed to marshal payload: %v", err)
	}
	return &event.EventPayload{
		Data: data,
		Kws:  &event.Websocket{}, // zero-value; do NOT assert on Emit output in these tests
	}
}

// ---------- Create_order ----------

func TestCreateOrder_Success(t *testing.T) {
	db := setupTestDB(t)
	table := seedTable(t, db)

	ep := newTestPayload(t, createOrderPayload{TableID: table.ID})
	Create_order(db, ep)

	var count int64
	db.Model(&models.Order{}).Where("table_id = ?", table.ID).Count(&count)
	if count != 1 {
		t.Errorf("expected 1 order created, got %d", count)
	}
}

func TestCreateOrder_InvalidTableID(t *testing.T) {
	db := setupTestDB(t)

	ep := newTestPayload(t, createOrderPayload{TableID: 9999}) // no such table
	Create_order(db, ep)

	var count int64
	db.Model(&models.Order{}).Count(&count)
	if count != 0 {
		t.Errorf("expected 0 orders after FK failure, got %d", count)
	}
}

func TestCreateOrder_InvalidJSON(t *testing.T) {
	db := setupTestDB(t)

	ep := &event.EventPayload{Data: []byte(`not json`), Kws: &event.Websocket{}}
	Create_order(db, ep) // should not panic

	var count int64
	db.Model(&models.Order{}).Count(&count)
	if count != 0 {
		t.Errorf("expected 0 orders on invalid JSON, got %d", count)
	}
}

// ---------- Order_food ----------

func TestOrderFood_Success(t *testing.T) {
	db := setupTestDB(t)
	table := seedTable(t, db)
	dish1 := seedDish(t, db, "Pad Thai")
	dish2 := seedDish(t, db, "Tom Yum")
	order := seedOrder(t, db, table.ID)

	payload := OrderFoodPayload{
		OrderID: order.ID,
		TableID: table.ID,
		OrderMap: map[uint]int{
			dish1.ID: 2,
			dish2.ID: 1,
		},
	}
	ep := newTestPayload(t, payload)
	Order_food(db, ep)

	var count int64
	db.Model(&models.Order_items{}).Where("order_id = ?", order.ID).Count(&count)
	if count != 2 {
		t.Errorf("expected 2 order items created, got %d", count)
	}
}

func TestOrderFood_EmptyOrderMap(t *testing.T) {
	db := setupTestDB(t)
	table := seedTable(t, db)
	order := seedOrder(t, db, table.ID)

	payload := OrderFoodPayload{OrderID: order.ID, TableID: table.ID, OrderMap: map[uint]int{}}
	ep := newTestPayload(t, payload)
	Order_food(db, ep)

	var count int64
	db.Model(&models.Order_items{}).Where("order_id = ?", order.ID).Count(&count)
	if count != 0 {
		t.Errorf("expected 0 order items for empty map, got %d", count)
	}
}

// This is the important one: one valid dish + one invalid dish in the same
// order. If the transaction is working correctly, EITHER both items exist
// or NEITHER does — never exactly one.
func TestOrderFood_PartialFailureRollsBackEverything(t *testing.T) {
	db := setupTestDB(t)
	table := seedTable(t, db)
	validDish := seedDish(t, db, "Pad Thai")
	const invalidDishID = uint(99999) // does not exist -> FK violation
	order := seedOrder(t, db, table.ID)

	payload := OrderFoodPayload{
		OrderID: order.ID,
		TableID: table.ID,
		OrderMap: map[uint]int{
			validDish.ID:  1,
			invalidDishID: 1,
		},
	}
	ep := newTestPayload(t, payload)
	Order_food(db, ep)

	var count int64
	db.Model(&models.Order_items{}).Where("order_id = ?", order.ID).Count(&count)
	if count != 0 {
		t.Errorf("expected 0 order items after rollback (all-or-nothing), got %d — "+
			"partial commit means the transaction isn't actually rolling back", count)
	}
}

func TestOrderFood_InvalidJSON(t *testing.T) {
	db := setupTestDB(t)

	ep := &event.EventPayload{Data: []byte(`not json`), Kws: &event.Websocket{}}
	Order_food(db, ep) // should not panic

	var count int64
	db.Model(&models.Order_items{}).Count(&count)
	if count != 0 {
		t.Errorf("expected 0 order items on invalid JSON, got %d", count)
	}
}

// ---------- Delete_order ----------

func TestDeleteOrder_Success(t *testing.T) {
	db := setupTestDB(t)
	table := seedTable(t, db)
	order := seedOrder(t, db, table.ID)

	payload := DeleteOrderPayload{Order: order}
	ep := newTestPayload(t, payload)
	Delete_order(db, ep)

	var count int64
	db.Unscoped().Model(&models.Order{}).Where("id = ?", order.ID).Count(&count)
	if count != 0 {
		t.Errorf("expected order to be hard-deleted, got %d matching rows", count)
	}
}
