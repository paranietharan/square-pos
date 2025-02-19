package pos

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"square-pos/pkg/config"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PosStore struct {
	db *gorm.DB
}

// NewStore initializes the user store with a database connection
func NewPosStore(db *gorm.DB) *PosStore {
	return &PosStore{db: db}
}

func (ps *PosStore) CreateOrder() CreateOrderResponse {
	idempotencyKey := uuid.New().String()
	orderReq := OrderRequest{
		IdempotencyKey: idempotencyKey,
		Order: Order{
			LocationID: config.GetEnv("LOCATION_ID", ""),
			LineItems: []LineItem{
				{
					Name:     "Book-01",
					Quantity: "100",
					BasePriceMoney: Money{
						Amount:   10,
						Currency: "USD",
					},
				},
			},
		},
	}

	jsonData, err := json.Marshal(orderReq)
	if err != nil {
		log.Fatalf("Error marshalling JSON: %v", err)
	}

	req, err := http.NewRequestWithContext(context.TODO(), "POST",
		config.GetEnv("CREATE_ORDER_API_URL", ""),
		bytes.NewBuffer(jsonData))

	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Authorization", "Bearer "+config.GetEnv("ACCESS_TOKEN", ""))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Square-Version", "2025-01-23")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	var orderResponse CreateOrderResponse
	if err := json.NewDecoder(resp.Body).Decode(&orderResponse); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	// Print the parsed response
	//fmt.Printf("Order Created Successfully: %+v\n", orderResponse)
	return orderResponse
}
