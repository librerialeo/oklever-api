package service

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken generate a token string for userId, rol with short or long expiration
func (s *Service) CreateToken(user int32, rol int32, long bool) (string, error) {
	claim := jwt.MapClaims{
		"authorized": true,
		"user":       user,
		"rol":        rol,
		"created":    time.Now(),
		"long":       long,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	return at.SignedString(([]byte)(os.Getenv("TOKEN_SECRET")))
}

// CheckToken validate token and return a new one if necessary
func (s *Service) CheckToken(tokenString string) (jwt.MapClaims, string, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ([]byte)(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, "", false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		created, createdOk := claims["created"].(string)
		userID, userIDOk := claims["user"].(int32)
		long, longOk := claims["long"].(bool)
		fmt.Println(createdOk, claims["created"], created)
		if createdOk && userIDOk && longOk {
			lastaction, err := s.GetUserLastAction(userID)
			if err != nil {
				fmt.Println("Token Data Corruption - Change token secret")
				return nil, "", false
			}
			actionTime := time.Now()

			if long {
				claims["created"] = actionTime
				at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
				newToken, err := at.SignedString(([]byte)(os.Getenv("TOKEN_SECRET")))
				if err != nil {
					return nil, "", true
				}
				return claims, newToken, true
			}
			createdTime, err := time.Parse(time.RFC3339Nano, created)
			if err != nil {
				return nil, "", true
			}
			durationInt, err := strconv.Atoi(os.Getenv("SHORT_SESSION"))
			if err != nil {
				return nil, "", true
			}
			if createdTime.Add(time.Duration(durationInt)).Unix() > lastaction.Unix() {
				claims["created"] = actionTime
				s.UpdateUserLastAction(userID, actionTime)
				at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
				newToken, err := at.SignedString(([]byte)(os.Getenv("TOKEN_SECRET")))
				if err != nil {
					return nil, "", true
				}
				return claims, newToken, true
			}
		}
		fmt.Println("Token Data Corruption - Change token secret")
		return nil, "", false
	}
	fmt.Println("Invalid Token")
	return nil, "", false
}
