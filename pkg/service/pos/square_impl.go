package pos

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"square-pos/pkg/config"
	"square-pos/pkg/dto"
	"strconv"

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

func (ps *PosStore) CreateOrder(request dto.CreateOrderRequest) dto.CreateOrderRes {
	productName := request.ProductName
	qunatity := request.Quantity
	amount := request.Amount

	idempotencyKey := uuid.New().String()
	qty := strconv.Itoa(qunatity)
	orderReq := dto.OrderRequest{
		IdempotencyKey: idempotencyKey,
		Order: dto.Order{
			LocationID: config.GetEnv("LOCATION_ID", ""),
			LineItems: []dto.LineItem{
				{
					Name:     productName,
					Quantity: qty,
					BasePriceMoney: dto.Money{
						Amount:   amount,
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

	var orderResponse dto.CreateOrderRes
	if err := json.NewDecoder(resp.Body).Decode(&orderResponse); err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	// Print the parsed response
	// fmt.Printf("Order Created Successfully: %+v\n", orderResponse)
	return orderResponse
}

func (ps *PosStore) GetOrder(orderID string) (*dto.CreateOrderRes, error) {
	url := fmt.Sprintf("https://connect.squareupsandbox.com/v2/orders/%s", orderID)

	req, err := http.NewRequestWithContext(context.TODO(), "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+config.GetEnv("ACCESS_TOKEN", ""))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Square-Version", "2025-01-23")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var orderResponse dto.CreateOrderRes
	if err := json.NewDecoder(resp.Body).Decode(&orderResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &orderResponse, nil
}

func (ps *PosStore) SubmitPayments(paymentReq dto.PaymentRequest) (*dto.PaymentResponse, error) {
	url := "https://connect.squareupsandbox.com/v2/payments"

	jsonData, err := json.Marshal(paymentReq)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payment request: %v", err)
	}

	req, err := http.NewRequestWithContext(context.TODO(), "POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+config.GetEnv("ACCESS_TOKEN", ""))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Square-Version", "2025-01-23")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var paymentResponse dto.PaymentResponse
	if err := json.NewDecoder(resp.Body).Decode(&paymentResponse); err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	return &paymentResponse, nil
}
