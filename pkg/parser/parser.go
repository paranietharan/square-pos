package parser

import (
	"log"
	"square-pos/pkg/dto"
	"strconv"

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

func ParseCreateOrderResponseToOrder(input *dto.CreateOrderRes) qlub.Order {
	log.Printf("Parsing order response: %+v", input)
	var order qlub.Order

	order.TableID = input.OrderRes.ReferenceID
	order.Hash = input.OrderRes.Id

	log.Printf("Set order.TableID: %s, order.Hash: %s, order.Status: %s", order.TableID, order.Hash, order.Status)
	for _, lineItem := range input.OrderRes.LineItems {
		log.Printf("Processing lineItem: %+v", lineItem)

		// Convert Quantity from string to int
		quantity, err := strconv.Atoi(lineItem.Quantity)
		if err != nil {
			log.Printf("Error converting quantity to int: %v", err)
			continue
		}
		log.Printf("Converted quantity: %d", quantity)

		finalPrice := lineItem.BasePriceMoney.Amount * float64(quantity)
		log.Printf("Calculated final price: %.2f", finalPrice)

		order.Items = append(order.Items, qlub.OrderItem{
			ID:         lineItem.UID,
			Title:      lineItem.Name,
			Quantity:   lineItem.Quantity,
			UnitPrice:  strconv.FormatFloat(lineItem.BasePriceMoney.Amount, 'f', -1, 64),
			BasePrice:  strconv.FormatFloat(lineItem.BasePriceMoney.Amount, 'f', -1, 64),
			FinalPrice: strconv.FormatFloat(finalPrice, 'f', -1, 64),
		})
	}

	log.Printf("Final order: %+v", order)
	return order
}

func ParsePaymentResponseToUpdatePaymentStatusCommand(input dto.PaymentResponse) qlub.UpdatePaymentStatusCommand {
	var updatePaymentStatusCmd qlub.UpdatePaymentStatusCommand

	updatePaymentStatusCmd.OrderID = input.Payment.ID
	updatePaymentStatusCmd.BillAmount = strconv.FormatFloat(input.Payment.AmountMoney.Amount, 'f', -1, 64)
	updatePaymentStatusCmd.TipAmount = strconv.FormatFloat(input.Payment.TotalMoney.Amount, 'f', -1, 64)
	updatePaymentStatusCmd.TotalAmount = strconv.FormatFloat(input.Payment.TotalMoney.Amount, 'f', -1, 64)

	return updatePaymentStatusCmd
}
