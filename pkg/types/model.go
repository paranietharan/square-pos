package types

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Email     string    `json:"email" gorm:"unique"`
	Password  string    `json:"-"` // Password is not exposed in API responses
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	Orders    []Order   `json:"orders"` // One-to-many relationship with orders
}

type Order struct {
	ID         int         `json:"id" gorm:"primaryKey"`
	UserID     int         `json:"userId"`
	User       User        `json:"user"`
	Total      float64     `json:"total"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	OrderItems []OrderItem `json:"orderItems"`
}

type OrderItem struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	OrderID   int     `json:"orderId"`
	Order     Order   `json:"order"`
	ProductID int     `json:"productId"`
	Product   Product `json:"product"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}

type Product struct {
	ID         int         `json:"id" gorm:"primaryKey"`
	Name       string      `json:"name"`
	Price      float64     `json:"price"`
	Stock      int         `json:"stock"`
	CreatedAt  time.Time   `json:"createdAt"`
	UpdatedAt  time.Time   `json:"updatedAt"`
	OrderItems []OrderItem `json:"orderItems"`
}
