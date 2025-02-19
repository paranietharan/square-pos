package pos

import (
	"net/http"
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
}

func (h *PosHandler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	res := h.store.CreateOrder()
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
