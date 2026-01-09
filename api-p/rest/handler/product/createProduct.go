package product

import (
	"encoding/json"
	"net/http"
	"server/repo"
	"server/utils"
)

func (p *Product) CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a product
	pd := repo.Product{}
	json.NewDecoder(r.Body).Decode(&pd)

	createdProduct, err := p.productRepo.Create(&pd)
	if err != nil {
		utils.ErrorData(w,
			map[string]string{"error": "Failed to create product"},
			http.StatusInternalServerError)
		return
	}

	pd = *createdProduct
	utils.SendData(w, pd, 201)
}
