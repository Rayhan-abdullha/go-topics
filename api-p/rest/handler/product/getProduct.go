package product

import (
	"net/http"
	"server/utils"
)

func (h *Handler) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	productList, err := h.svc.List()
	if err != nil {
		utils.ErrorData(w, map[string]string{"error": "Failed to retrieve products"}, http.StatusInternalServerError)
		return
	}
	utils.SendData(w, productList, 200)
}
