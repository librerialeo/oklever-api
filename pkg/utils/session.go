package utils

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken generate a token string for userId, rol with short or long expiration
func CreateToken(user int32, rol int32, long bool) (string, error) {
	claim := jwt.MapClaims{
		"authorized": true,
		"user": user,
		"rol": rol,
	}
	if long {
		claim["exp"] = os.Getenv("LOGN_SESSION")
	} else {
		claim["exp"] = os.Getenv("SHORT_SESSION")
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	return at.SignedString(([]byte)(os.Getenv("TOKEN_SECRET")))
}