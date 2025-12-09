package utils

import "net/http"

type HandlerSignature func(w http.ResponseWriter, r *http.Request)

func Handler(handler HandlerSignature) http.Handler {
	return http.HandlerFunc(handler)
}
