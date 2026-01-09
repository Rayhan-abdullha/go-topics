package utils

import (
	"crypto/hmac"
	"crypto/sha256"
)

func HmacSha256(data, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(data))
	return Base64UrlEncode(h.Sum(nil))
}
