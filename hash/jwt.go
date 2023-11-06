package hash

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func NewJWT(secret []byte, values map[string]interface{}) (string, error) {
	claims := jwt.MapClaims{"exp": time.Now().Add(time.Hour * 24)}
	for k, v := range values {
		claims[k] = v
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secret)
}
