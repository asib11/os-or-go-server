package middleware

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"net/http"
	"strings"
)

func (m *Middlewares) AuthenticateJWT(next http.Handler) http.Handler {
	// parse jwt
	// parse header and payload or claims
	// hmac-sha-256 algorithm -> hash hmac(header, payload, secret key)
	// parse signature part from the jwt
	// if the signature and hash is same => forward to create products
	// otherwise 401 status code with Unauthorized

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		headerArr := strings.Split(token, " ")
		if len(headerArr) != 2 || headerArr[0] != "Bearer" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		accessToken := headerArr[1]

		tokenParts := strings.Split(accessToken, ".")
		if len(tokenParts) != 3 {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		jwtHeader := tokenParts[0]
		jwtPayload := tokenParts[1]
		jwtSignature := tokenParts[2]

		if jwtHeader == "" || jwtPayload == "" || jwtSignature == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		message := jwtHeader + "." + jwtPayload

		byteArrSecret := []byte(m.conf.JWTSecret)
		byteArrMessage := []byte(message)

		h := hmac.New(sha256.New, byteArrSecret)
		h.Write(byteArrMessage)

		hash := h.Sum(nil)
		signature := base64UrlEncode(hash)

		if signature != jwtSignature {
			http.Error(w, "Unauthorized, you are hacker", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func base64UrlEncode(data []byte) string {
	return base64.URLEncoding.WithPadding(base64.NoPadding).EncodeToString(data)
}
