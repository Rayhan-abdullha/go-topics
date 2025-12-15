package user

import (
	middleware "article/rest/middlewares"
	"article/utils"
	"net/http"
)

func (h *Handlers) RegisterRoutes(mux *http.ServeMux, m *middleware.MiddleWareManager) {
	mux.Handle("GET /users", m.With(
		utils.Handler(h.GetUserHandler),
	))
	mux.Handle("POST /users", m.With(
		http.HandlerFunc(h.CreateUserHandler),
	))
}
