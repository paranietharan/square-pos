package pos

import (
	"encoding/json"
	"net/http"
	"square-pos/pkg/dto"
	"square-pos/pkg/types"
	"square-pos/pkg/utils"

	"github.com/gorilla/mux"
)

type PosHandler struct {
	store types.PosStore
}

func NewPosHandler(store types.PosStore) *PosHandler {
	return &PosHandler{store: store}
}

func (h *PosHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create-order", h.handleCreateOrder).Methods("POST")
	router.HandleFunc("/order/{id}", h.handleGetOrder).Methods("GET")
	router.HandleFunc("/submit-payment", h.handleSubmitPayment).Methods("POST")
}

func (h *PosHandler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var createOrderReq dto.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&createOrderReq); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}
	res := h.store.CreateOrder(createOrderReq)
	utils.WriteJSON(w, http.StatusCreated, res)
}

func (h *PosHandler) handleGetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	res, _ := h.store.GetOrder(orderID)

	if res == nil {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "Order not found",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, res)
}

func (h *PosHandler) handleSubmitPayment(w http.ResponseWriter, r *http.Request) {
	var paymentReq dto.PaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&paymentReq); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}

	paymentResp, err := h.store.SubmitPayments(paymentReq)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, paymentResp)
}
