package main

import (
	"fmt"
	"net/http"
	"server/db"
	"server/handler"
	mid "server/middleware"
)

var (
	count = 0
)

func main() {
	mux := http.NewServeMux()
	mng := mid.NewManager()
	mng.Use(
		mid.LoggerMid,
		mid.PreflightReq,
		mid.Cors,
	)

	mux.Handle("GET /products", mng.With(
		http.HandlerFunc(handler.GetProductHandler),
	))
	mux.Handle("POST /products", mng.With(
		http.HandlerFunc(handler.CreateProductHandler),
		mid.AuthMid,
	))
	// mux.Handle("PUT /product", mng.With(
	// 	http.HandlerFunc(handler.UpdateProductHandler),
	// 	mid.AuthMid,
	// ))
	// mux.Handle("DELETE /product", mng.With(
	// 	http.HandlerFunc(handler.DeleteProductHandler),
	// 	mid.AuthMid,
	// ))
	muxWrapper := mng.WrapMux(mux)
	fmt.Println("Server is running on port 4000")
	http.ListenAndServe(":4000", muxWrapper)
}

func init() {
	db.ProductList = []db.Product{
		{ID: 1, Title: "Product 1", Description: "Description 1", Price: 100, ImgUrl: "https://example.com/image1.jpg"},
		{ID: 2, Title: "Product 2", Description: "Description 2", Price: 200, ImgUrl: "https://example.com/image2.jpg"},
	}
}
