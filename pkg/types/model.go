package types

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	Orders    []Order   `json:"orders"`
}

type Order struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	UserID      int       `json:"userId"`
	User        User      `json:"user"`
	LocationID  string    `json:"location_id"`
	OrderID     string    `json:"order_id"`
	ProductName string    `json:"product_name"`
	Total       float64   `json:"total"`
	Quantity    int       `json:"quantity"`
	UnitPrice   float64   `json:"unitPrice"`
	IsPaid      bool      `json:"is_paid"`
	TableID     string    `json:"table_id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func (o *Order) CalculateTotal() {
	o.Total = float64(o.Quantity) * o.UnitPrice
}
