package product

import (
	"net/http"
	"server/rest/middleware"
)

func (p *Product) RegisterRoutes(mux *http.ServeMux, mng *middleware.Middlewares) {
	mux.Handle("GET /products", mng.With(
		http.HandlerFunc(p.GetProductHandler),
	))
	mux.Handle("POST /products", mng.With(
		http.HandlerFunc(p.CreateProductHandler),
		p.middlewares.AuthMid,
	))
}
