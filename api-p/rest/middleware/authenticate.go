package middleware

import (
	"crypto/hmac"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"server/config"
	"server/utils"
	"strings"
)

func VerifyJWT(token, secret string) (map[string]interface{}, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, errors.New("invalid token format")
	}

	unsigned := parts[0] + "." + parts[1]
	expectedSig := utils.HmacSha256(unsigned, secret)

	if !hmac.Equal([]byte(expectedSig), []byte(parts[2])) {
		return nil, errors.New("invalid signature")
	}

	payloadBytes, err := utils.Base64UrlDecode(parts[1])
	if err != nil {
		return nil, err
	}

	var payload map[string]interface{}
	json.Unmarshal(payloadBytes, &payload)

	return payload, nil
}

func (m *MiddlewareType) AuthMid(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			utils.ErrorData(w, map[string]string{"error": "Unauthorized"}, http.StatusUnauthorized)
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			utils.ErrorData(w, map[string]string{"error": "Invalid token format"}, http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		secret := config.GetConfig().SecretJWT

		_, err := VerifyJWT(token, secret)
		if err != nil {
			fmt.Println("Token verification error:", err)
			utils.ErrorData(w, map[string]string{"error": err.Error()}, http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}
