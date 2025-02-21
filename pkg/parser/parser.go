package parser

import (
	"log"
	"square-pos/pkg/dto"

	"github.com/clubpay-pos-worker/sdk-go/v2/qlub"
)

func ParseOrderInputToOrderRequest(input qlub.OrderInput) dto.OrderRequest {
	log.Printf("\nQlub order Input :\n%v\n", input)
	var orderReq dto.OrderRequest

	orderReq.IdempotencyKey = input.ID

	// ----------------------------------------------
	// ----------------------- ORDER ----------------
	// ----------------------------------------------
	orderReq.Order.ReferenceID = input.TableID      // reference
	orderReq.Order.LocationID = input.RevenueCenter // assign revenue center as a Location id
	// loop through all products and append it
	for _, v := range input.Products {
		basePriceAmount := v.Qty.Value()
		orderReq.Order.LineItems = append(orderReq.Order.LineItems, dto.LineItem{
			Name:     v.ProductName,
			Quantity: v.Qty.String(),
			BasePriceMoney: dto.Money{
				Amount:   basePriceAmount,
				Currency: "USD",
			},
		})
	}

	log.Printf("Converted orderReq is\n%v\n---------------------------------------", orderReq)

	return orderReq
}
