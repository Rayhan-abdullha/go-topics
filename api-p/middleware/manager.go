package middleware

import "net/http"

type Middleware func(next http.Handler) http.Handler
type Middlewares struct {
	middlewares []Middleware
}

func NewManager() *Middlewares {
	return &Middlewares{
		middlewares: make([]Middleware, 0),
	}
}
func (mng *Middlewares) Use(mid ...Middleware) {
	mng.middlewares = append(mng.middlewares, mid...)
}
func (mng *Middlewares) With(next http.Handler, mid ...Middleware) http.Handler {

	for _, middleware := range mid {
		next = middleware(next)
	}

	return next
}

func (mng *Middlewares) WrapMux(mux http.Handler, mid ...Middleware) http.Handler {

	// global middleware
	for _, middleware := range mng.middlewares {
		mux = middleware(mux)
	}

	return mux
}
