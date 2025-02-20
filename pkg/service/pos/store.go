package pos

import (
	"square-pos/pkg/types"

	"gorm.io/gorm"
)

// function to store order in the db
func CreateOrder(userID int, quantity int, unitPrice float64, db *gorm.DB) error {
	order := &types.Order{
		UserID:    userID,
		Quantity:  quantity,
		UnitPrice: unitPrice,
	}

	order.CalculateTotal()

	result := db.Create(order)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
