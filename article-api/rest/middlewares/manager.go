package middlewares

import "net/http"

type Middlewares func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middlewares
}

func NewManager() *Manager {
	return &Manager{}
}

func (m *Manager) Use(mid ...Middlewares) {
	m.globalMiddlewares = append(m.globalMiddlewares, mid...)
}

func (m *Manager) WrapMux(next http.Handler) http.Handler {
	for _, m := range m.globalMiddlewares {
		next = m(next)
	}
	return next
}

func (m *Manager) With(next http.Handler, mid ...Middlewares) http.Handler {
	for _, m := range mid {
		next = m(next)
	}
	return next
}
