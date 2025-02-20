package pos

import (
	"log"
	"square-pos/pkg/types"
	"time"

	"gorm.io/gorm"
)

// function to store order in the db
func CreateOrder(user types.User, LocationID string, orderID string, productName string, qty int, up float64, tableID string, db *gorm.DB) error {
	log.Printf("User that create order : %v", user)

	order := types.Order{
		UserID:      user.ID,
		LocationID:  LocationID,
		OrderID:     orderID,
		User:        user,
		ProductName: productName,
		Quantity:    qty,
		UnitPrice:   up,
		TableID:     tableID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	order.CalculateTotal()

	result := db.Create(&order)
	if result.Error != nil {
		return result.Error
	}

	log.Printf("Created order : %v", order)
	return nil
}

func UpdatePaymentsInDB(OrderId string, LocationID string, db *gorm.DB) error {
	result := db.Exec("UPDATE orders SET is_paid = ?, updated_at = ? WHERE order_id = ? AND location_id = ?", true, time.Now(), OrderId, LocationID)

	if result.Error != nil {
		log.Printf("Error updating payment status: %v", result.Error)
		return result.Error
	}

	if result.RowsAffected == 0 {
		log.Printf("Order not found: LocationID = %s, OrderID = %s", LocationID, OrderId)
		return nil
	}

	log.Printf("Updated payment status for OrderID = %s at LocationID = %s", OrderId, LocationID)
	return nil
}

func GetOrdersByTableID(tableID string, db *gorm.DB) ([]types.Order, error) {
	var orders []types.Order

	result := db.Where("table_id = ?", tableID).Find(&orders)
	if result.Error != nil {
		log.Printf("Error retrieving orders for table %s: %v", tableID, result.Error)
		return nil, result.Error
	}

	log.Printf("Retrieved orders for table %s: %v", tableID, orders)
	return orders, nil
}
