package product

import (
	"net/http"
	"server/rest/middleware"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, mng *middleware.Middlewares) {
	mux.Handle("GET /products", mng.With(
		http.HandlerFunc(h.GetProductHandler),
	))
	mux.Handle("POST /products", mng.With(
		http.HandlerFunc(h.CreateProductHandler),
		h.middlewares.AuthMid,
	))
}
