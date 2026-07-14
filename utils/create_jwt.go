package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload struct {
	Sub         string `json:"sub"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func CreateJWT(secret string, payload Payload) (string, error) {
	header := Header{
		Alg: "HS256",
		Typ: "JWT",
	}

	bytesHeader, err := json.Marshal(header)
	if err != nil {
		return "", err
	}

	headerB64 := base64UrlEncode(bytesHeader)

	bytesData, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	payloadB64 := base64UrlEncode(bytesData)

	message := headerB64 + "." + payloadB64

	byteArrSecret := []byte(secret)
	byteArrMessage := []byte(message)

	h := hmac.New(sha256.New, byteArrSecret)
	h.Write(byteArrMessage)

	signature := h.Sum(nil)
	signatureB64 := base64UrlEncode(signature)

	jwt := message + "." + signatureB64
	
	return jwt, nil
}


func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
