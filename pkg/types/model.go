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
	ID        int       `json:"id" gorm:"primaryKey"`
	UserID    int       `json:"userId"`
	User      User      `json:"user"`
	Total     float64   `json:"total"`
	Quantity  int       `json:"quantity"`
	UnitPrice float64   `json:"unitPrice"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (o *Order) CalculateTotal() {
	o.Total = float64(o.Quantity) * o.UnitPrice
}
