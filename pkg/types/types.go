package types

import (
	"square-pos/pkg/dto"
)

type RegisterUserPayload struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=3,max=130"`
}

type LoginUserPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserStore interface {
	GetUserByEmail(email string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(*User) error
}

type PosStore interface {
	CreateOrder(dto.CreateOrderRequest, User) dto.CreateOrderRes
	GetOrder(orderID string) (*dto.CreateOrderRes, error)
	SubmitPayments(paymentReq dto.PaymentRequest) (*dto.PaymentResponse, error)
}
