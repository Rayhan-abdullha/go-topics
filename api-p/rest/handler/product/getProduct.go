package product

import (
	"net/http"
	"server/utils"
)

func (p *Product) GetProductHandler(w http.ResponseWriter, r *http.Request) {
	productList, err := p.productRepo.List()
	if err != nil {
		utils.ErrorData(w, map[string]string{"error": "Failed to retrieve products"}, http.StatusInternalServerError)
		return
	}
	utils.SendData(w, productList, 200)
}
