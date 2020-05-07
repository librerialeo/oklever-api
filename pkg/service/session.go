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
		fmt.Println("a")
		return nil, "", false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		created, createdOk := claims["created"]
		user, userOk := claims["user"]
		long, longOk := claims["long"]
		if createdOk && userOk && longOk {
			actionTime := time.Now()
			if long.(bool) {
				s.UpdateUserLastAction(int32(user.(float64)), actionTime)
				fmt.Println("c")
				return claims, "", true
			}
			lastaction, err := s.GetUserLastAction(int32(user.(float64)))
			if err != nil {
				fmt.Println("can't get last action", err)
				fmt.Println("b")
				return nil, "", true
			}
			createdTime, err := time.Parse(time.RFC3339Nano, created.(string))
			if err != nil {
				fmt.Println("d")
				return nil, "", true
			}
			durationInt, err := strconv.Atoi(os.Getenv("SHORT_SESSION"))
			if err != nil {
				fmt.Println("e")
				return nil, "", true
			}
			expDuration := time.Duration(durationInt)
			expirationTime := createdTime.Add(expDuration)
			fmt.Println("exp:", expirationTime, "\nlast:", lastaction, "\nnow:", actionTime)
			if expirationTime.Unix() < actionTime.Unix() {
				if lastaction.Add(expDuration).Unix() > actionTime.Unix() {
					claims["created"] = actionTime
					at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
					newToken, err := at.SignedString(([]byte)(os.Getenv("TOKEN_SECRET")))
					if err != nil {
						newLast, err := s.UpdateUserLastAction(int32(user.(float64)), actionTime)
						if err != nil {
							fmt.Println("new:", newLast)
						}
						fmt.Println("f")
						return claims, "", true
					}
					newLast, err := s.UpdateUserLastAction(int32(user.(float64)), actionTime)
					if err != nil {
						fmt.Println("new:", newLast)
					}
					fmt.Println("g")
					return claims, newToken, true
				}
			} else {
				newLast, err := s.UpdateUserLastAction(int32(user.(float64)), actionTime)
				if err != nil {
					fmt.Println("new:", newLast)
				}
				fmt.Println("h")
				return claims, "", true
			}
			fmt.Println("i")
			return nil, "", true
		}
		fmt.Println("Token Data Corruption - Change token secret")
		fmt.Println("j")
		return nil, "", false
	}
	fmt.Println("Invalid Token")
	fmt.Println("k")
	return nil, "", false
}
