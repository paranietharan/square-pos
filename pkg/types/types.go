package types

import (
	"square-pos/pkg/dto"

	"github.com/clubpay-pos-worker/sdk-go/v2/qlub"
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
	CreateUser(*User) (dto.UserCreateResponse, error)
}

type PosStore interface {
	CreateOrder(qlub.OrderInput, User) error
	GetOrder(orderID string) (order qlub.Order, err error)
	SubmitPayments(paymentReq dto.PaymentRequest) (*qlub.UpdatePaymentStatusCommand, error)
	//GetOrdersByTableID(tableID string) ([]*dto.CreateOrderRes, error)
}
