package utils

import (
	"encoding/base64"
	"strings"
)

func Base64UrlDecode(data string) ([]byte, error) {
	if m := len(data) % 4; m != 0 {
		data += strings.Repeat("=", 4-m)
	}
	return base64.URLEncoding.DecodeString(data)
}
