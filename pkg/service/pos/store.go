package pos

import (
	"log"
	"square-pos/pkg/types"
	"time"

	"gorm.io/gorm"
)

// function to store order in the db
func CreateOrder(user types.User, LocationID string, orderID string, productName string, qty int, up float64, db *gorm.DB) error {
	log.Printf("User that create order : %v", user)

	order := types.Order{
		UserID:      user.ID,
		LocationID:  LocationID,
		OrderID:     orderID,
		User:        user,
		ProductName: productName,
		Quantity:    qty,
		UnitPrice:   up,
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
