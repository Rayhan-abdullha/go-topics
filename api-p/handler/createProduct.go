package handler

import (
	"encoding/json"
	"net/http"
	"server/db"
	"server/utils"
)

func CreateProductHandler(w http.ResponseWriter, r *http.Request) {
	// Implementation for creating a product
	pd := db.Product{}
	json.NewDecoder(r.Body).Decode(&pd)

	db.ProductList = append(db.ProductList, pd)

	utils.SendData(w, pd, 201)
}
