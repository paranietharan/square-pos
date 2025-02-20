package dto

import (
	"strconv"
	"time"
)

type Money struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type LineItem struct {
	Name           string `json:"name"`
	Quantity       string `json:"quantity"`
	BasePriceMoney Money  `json:"base_price_money"`
}

type Order struct {
	LocationID string     `json:"location_id"`
	LineItems  []LineItem `json:"line_items"`
}

type OrderRequest struct {
	IdempotencyKey string `json:"idempotency_key"`
	Order          Order  `json:"order"`
}

type NetAmounts struct {
	TotalMoney         Money `json:"total_money"`
	TaxMoney           Money `json:"tax_money"`
	DiscountMoney      Money `json:"discount_money"`
	TipMoney           Money `json:"tip_money"`
	ServiceChargeMoney Money `json:"service_charge_money"`
}

type Source struct {
	Name string `json:"name"`
}

type LineItemRes struct {
	UID                      string `json:"uid"`
	Quantity                 string `json:"quantity"`
	Name                     string `json:"name"`
	BasePriceMoney           Money  `json:"base_price_money"`
	GrossSalesMoney          Money  `json:"gross_sales_money"`
	TotalTaxMoney            Money  `json:"total_tax_money"`
	TotalDiscountMoney       Money  `json:"total_discount_money"`
	TotalMoney               Money  `json:"total_money"`
	VariationTotalPriceMoney Money  `json:"variation_total_price_money"`
	ItemType                 string `json:"item_type"`
	TotalServiceChargeMoney  Money  `json:"total_service_charge_money"`
}

type CreateOrderResponse struct {
	Id                      string        `json:"id"`
	LocationID              string        `json:"location_id"`
	LineItems               []LineItemRes `json:"line_items"`
	CreatedAt               string        `json:"created_at"`
	UpdatedAt               string        `json:"updated_at"`
	State                   string        `json:"state"`
	Version                 int           `json:"version"`
	TotalTaxMoney           Money         `json:"total_tax_money"`
	TotalDiscountMoney      Money         `json:"total_discount_money"`
	TotalTipMoney           Money         `json:"total_tip_money"`
	TotalMoney              Money         `json:"total_money"`
	TotalServiceChargeMoney Money         `json:"total_service_charge_money"`
	NetAmounts              NetAmounts    `json:"net_amounts"`
	Source                  Source        `json:"source"`
	NetAmountDueMoney       Money         `json:"net_amount_due_money"`
}

type CreateOrderRes struct { // used to receive the request from the pos
	OrderRes CreateOrderResponse `json:"order"`
}

// dtos for payment
type PaymentRequest struct {
	IdempotencyKey    string      `json:"idempotency_key"`
	AmountMoney       Money       `json:"amount_money"`
	SourceID          string      `json:"source_id"`
	OrderID           string      `json:"order_id"`
	AcceptPartialAuth bool        `json:"accept_partial_authorization"`
	LocationID        string      `json:"location_id"`
	ReferenceID       string      `json:"reference_id"`
	CashDetails       CashDetails `json:"cash_details"`
}

type CashDetails struct {
	BuyerSuppliedMoney Money `json:"buyer_supplied_money"`
}

type PaymentResponse struct {
	Payment struct {
		ID            string      `json:"id"`
		AmountMoney   Money       `json:"amount_money"`
		Status        string      `json:"status"`
		SourceType    string      `json:"source_type"`
		LocationID    string      `json:"location_id"`
		OrderID       string      `json:"order_id"`
		ReferenceID   string      `json:"reference_id"`
		TotalMoney    Money       `json:"total_money"`
		CashDetails   CashDetails `json:"cash_details"`
		ReceiptNumber string      `json:"receipt_number"`
		ReceiptURL    string      `json:"receipt_url"`
	} `json:"payment"`
}

// dto for create order
type CreateOrderRequest struct {
	ProductName string  `json:"product_name"`
	Quantity    int     `json:"quantity"`
	Amount      float64 `json:"amount"`
}

// Dto for sending to the user - Order Response
type CreateOrderResp struct {
	ID       string      `json:"order_id"`
	OpenedAt time.Time   `json:"opened_at"`
	IsClosed bool        `json:"is_closed"`
	Table    string      `json:"table"`
	Items    []OrderItem `json:"items"`
	Totals   OrderTotals `json:"totals"`
}

type OrderItem struct {
	Name      string          `json:"name"`
	Comment   string          `json:"comment"`
	UnitPrice float64         `json:"unit_price"`
	Quantity  int             `json:"quantity"`
	Discounts []OrderDiscount `json:"discounts"`
	Modifiers []OrderModifier `json:"modifiers"`
	Amount    float64         `json:"amount"`
}

type OrderDiscount struct {
	Name         string  `json:"name"`
	IsPercentage bool    `json:"is_percentage"`
	Value        float64 `json:"value"`
	Amount       float64 `json:"amount"`
}

type OrderModifier struct {
	Name      string  `json:"name"`
	UnitPrice float64 `json:"unit_price"`
	Quantity  int     `json:"quantity"`
	Amount    float64 `json:"amount"`
}

type OrderTotals struct {
	Discounts     float64 `json:"discounts"`
	Due           float64 `json:"due"`
	Tax           float64 `json:"tax"`
	ServiceCharge float64 `json:"service_charge"`
	Paid          float64 `json:"paid"`
	Tips          float64 `json:"tips"`
	Total         float64 `json:"total"`
}

func ParseCreateOrderResponse(response CreateOrderResponse) CreateOrderResp {
	orderResp := CreateOrderResp{
		ID:       response.Id,
		OpenedAt: parseTime(response.CreatedAt),
		IsClosed: response.State == "CLOSED", // Default set the value as order cloased
		Table:    "29",                       // fixed table
		Items:    parseItems(response.LineItems),
		Totals:   parseTotals(response.NetAmounts),
	}

	return orderResp
}

func parseTime(timeStr string) time.Time {
	parsedTime, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}
	}
	return parsedTime
}

func parseItems(lineItems []LineItemRes) []OrderItem {
	var items []OrderItem
	for _, item := range lineItems {
		discounts := parseDiscounts(item.TotalDiscountMoney)
		modifiers := parseModifiers(item.TotalServiceChargeMoney)

		orderItem := OrderItem{
			Name:      item.Name,
			Comment:   "",
			UnitPrice: item.BasePriceMoney.Amount,
			Quantity:  parseQuantity(item.Quantity),
			Discounts: discounts,
			Modifiers: modifiers,
			Amount:    item.TotalMoney.Amount,
		}
		items = append(items, orderItem)
	}
	return items
}

func parseDiscounts(discountMoney Money) []OrderDiscount {
	return []OrderDiscount{
		{
			Name:         "",
			IsPercentage: true,
			Value:        discountMoney.Amount,
			Amount:       discountMoney.Amount,
		},
	}
}

func parseModifiers(serviceChargeMoney Money) []OrderModifier {
	return []OrderModifier{
		{
			Name:      "",
			UnitPrice: serviceChargeMoney.Amount,
			Quantity:  1,
			Amount:    serviceChargeMoney.Amount,
		},
	}
}

func parseQuantity(quantityStr string) int {
	quantity, err := strconv.Atoi(quantityStr)
	if err != nil {
		return 0
	}
	return quantity
}

func parseTotals(netAmounts NetAmounts) OrderTotals {
	return OrderTotals{
		Discounts:     netAmounts.DiscountMoney.Amount,
		Due:           netAmounts.TotalMoney.Amount,
		Tax:           netAmounts.TaxMoney.Amount,
		ServiceCharge: netAmounts.ServiceChargeMoney.Amount,
		Paid:          netAmounts.TotalMoney.Amount,
		Tips:          netAmounts.TipMoney.Amount,
		Total:         netAmounts.TotalMoney.Amount,
	}
}
