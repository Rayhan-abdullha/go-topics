package utils

import (
	"encoding/base64"
	"strings"
)

func Base64UrlEncode(data []byte) string {
	return strings.TrimRight(
		base64.URLEncoding.EncodeToString(data),
		"=",
	)
}
