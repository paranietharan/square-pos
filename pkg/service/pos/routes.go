package pos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"square-pos/pkg/dto"
	"square-pos/pkg/service/auth"
	"square-pos/pkg/types"
	"square-pos/pkg/utils"

	"github.com/gorilla/mux"
)

type PosHandler struct {
	posStore  types.PosStore
	userStore types.UserStore
}

func NewPosHandler(posStore types.PosStore, userStore types.UserStore) *PosHandler {
	return &PosHandler{
		posStore:  posStore,
		userStore: userStore,
	}
}

func (h *PosHandler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/create-order", auth.WithJWTAuth(h.handleCreateOrder, h.userStore)).Methods("POST")
	router.HandleFunc("/order/{id}", auth.WithJWTAuth(h.handleGetOrder, h.userStore)).Methods("GET")
	router.HandleFunc("/order/table/{id}", auth.WithJWTAuth(h.handleGetOrderByTable, h.userStore)).Methods("GET")
	router.HandleFunc("/submit-payment", auth.WithJWTAuth(h.handleSubmitPayment, h.userStore)).Methods("POST")
}

func (h *PosHandler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {

	// get user info ---------------------------
	userID := auth.GetUserIDFromContext(r.Context())
	if userID == -1 {
		utils.WriteError(w, http.StatusUnauthorized, fmt.Errorf("invalid or missing user ID"))
		return
	}

	user, err := h.userStore.GetUserByID(userID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to fetch user: %v", err))
		return
	}
	// ----------------------------------

	var createOrderReq dto.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&createOrderReq); err != nil {
		utils.WriteJSON(w, http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
		return
	}
	res := h.posStore.CreateOrder(createOrderReq, *user)
	utils.WriteJSON(w, http.StatusCreated, res)
}

func (h *PosHandler) handleGetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	res, _ := h.posStore.GetOrder(orderID)

	if res == nil {
		utils.WriteJSON(w, http.StatusNotFound, map[string]string{
			"error": "Order not found",
		})
		return
	}

	utils.WriteJSON(w, http.StatusOK, res)
}

func (h *PosHandler) handleGetOrderByTable(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	orderID := vars["id"]

	res, _ := h.posStore.GetOrdersByTableID(orderID)

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

	paymentResp, err := h.posStore.SubmitPayments(paymentReq)
	if err != nil {
		utils.WriteJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	utils.WriteJSON(w, http.StatusOK, paymentResp)
}
