package product

import (
	"encoding/json"
	"net/http"
	"server/domain"
	"server/utils"
)

func (h *Handler) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a product
	pd := domain.Product{}
	json.NewDecoder(r.Body).Decode(&pd)

	createdProduct, err := h.svc.Create(&pd)
	if err != nil {
		utils.ErrorData(w,
			map[string]string{"error": "Failed to create product"},
			http.StatusInternalServerError)
		return
	}

	pd = *createdProduct
	utils.SendData(w, pd, 201)
}
