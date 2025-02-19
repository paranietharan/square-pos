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
}

func (h *PosHandler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	h.store.CreateOrder()
	utils.WriteJSON(w, http.StatusCreated, nil)
}
