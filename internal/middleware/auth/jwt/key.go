package jwt

import (
	"github.com/golang-jwt/jwt/v4"
)

// yourKeyFunc is a key function that provides the JWT signing key.
func yourKeyFunc(token *jwt.Token) (interface{}, error) {
	switch token.Method.Alg() {
	case jwt.SigningMethodHS256.Alg():
		return []byte("your_secret_key"), nil // Your HMAC secret key
	default:
		return nil, ErrUnSupportSigningMethod // Unsupported signing method
	}
}
