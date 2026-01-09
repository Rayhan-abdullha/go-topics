package utils

import (
	"encoding/json"
)

type Header struct {
	Algo string `json:"algo"`
	Typ  string `json:"typ"`
}
type Payload struct {
	Sub   int    `json:"sub"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func CreateJWT(payload Payload, secret string) (string, error) {
	header := map[string]string{
		"alg": "HS256",
		"typ": "JWT",
	}

	hJSON, _ := json.Marshal(header)
	pJSON, _ := json.Marshal(payload)

	headerEnc := Base64UrlEncode(hJSON)
	payloadEnc := Base64UrlEncode(pJSON)

	unsignedToken := headerEnc + "." + payloadEnc
	signature := HmacSha256(unsignedToken, secret)

	return unsignedToken + "." + signature, nil
}
