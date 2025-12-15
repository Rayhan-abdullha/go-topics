package user

import (
	"article/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, m *middlewares.Manager) {
	mux.Handle("GET /users", m.With(
		http.HandlerFunc(h.GetUser),
		h.middleware.Authenticate,
	))
}
