package handler

import (
	"net/http"
	"server/db"
	"server/utils"
)

func GetProductHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	utils.SendData(w, db.ProductList, 200)
}
