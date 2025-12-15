package user

import (
	"encoding/json"
	"net/http"
)

func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("hello user")

}
