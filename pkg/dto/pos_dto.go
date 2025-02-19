package dto

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

type CreateOrderRes struct {
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
