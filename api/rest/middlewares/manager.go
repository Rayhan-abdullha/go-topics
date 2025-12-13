package middleware

import (
	"net/http"
)

type MiddleWareManager struct {
	middlewares []Middlewares
}

func NewManager() *MiddleWareManager {
	return &MiddleWareManager{
		middlewares: make([]Middlewares, 0),
	}
}

func (m *MiddleWareManager) Use(middlewares ...Middlewares) {
	m.middlewares = append(m.middlewares, middlewares...)
}

type Middlewares func(http.Handler) http.Handler

func (m *MiddleWareManager) With(next http.Handler, mid ...Middlewares) http.Handler {

	// apply middlewares
	for _, m := range mid {
		next = m(next)
	}
	return next
}

func (m *MiddleWareManager) WrapMux(next http.Handler) http.Handler {
	for _, m := range m.middlewares {
		next = m(next)
	}
	return next
}
