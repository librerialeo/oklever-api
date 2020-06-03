package service

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// CreateToken generate a token string for userId, rol with short or long expiration
func (s *Service) CreateToken(user int32, rol int32, status string, long bool) (string, error) {
	claim := jwt.MapClaims{
		"authorized": true,
		"user":       user,
		"rol":        rol,
		"status":     status,
		"created":    time.Now(),
		"long":       long,
	}
	at := jwt.NewWithClaims(jwt.SigningMethodHS512, claim)
	return at.SignedString(([]byte)(os.Getenv("TOKEN_SECRET")))
}

// CheckToken validate token and return a new one if necessary
func (s *Service) CheckToken(tokenString string) (*UserCredential, string, bool) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return ([]byte)(os.Getenv("TOKEN_SECRET")), nil
	})
	if err != nil {
		return nil, "", false
	}
	if !token.Valid {
		fmt.Println("Invalid Token")
		return nil, "", false
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("Invalid Claim")
		return nil, "", false
	}
	created, createdOk := claims["created"].(string)
	user, userOk := claims["user"].(float64)
	rol, rolOk := claims["rol"].(float64)
	status, statusOk := claims["status"].(string)
	long, longOk := claims["long"].(bool)
	if !createdOk || !userOk || !longOk || !rolOk || !statusOk {
		fmt.Println("Token Data Corruption - Change token secret")
		return nil, "", false
	}
	credential, err := s.GetUserCredentialsByID(int32(user))
	if err != nil || credential == nil {
		fmt.Println("Token validation: no credential found")
		return nil, "", false
	}
	if credential.Rol != int32(rol) || credential.Status != status {
		fmt.Println("Token validation: invalid credential")
		return nil, "", false
	}
	actionTime := time.Now()
	if long {
		s.UpdateUserLastAction(int32(user), actionTime)
		return credential, "", true
	}
	lastaction, err := s.GetUserLastAction(int32(user))
	if err != nil {
		fmt.Println("can't get last action", err)
		return nil, "", true
	}
	createdTime, err := time.Parse(time.RFC3339Nano, created)
	if err != nil {
		return nil, "", true
	}
	durationInt, err := strconv.Atoi(os.Getenv("SHORT_SESSION"))
	if err != nil {
		return nil, "", true
	}
	expDuration := time.Duration(time.Second.Nanoseconds() * int64(durationInt))
	expirationTime := createdTime.Add(expDuration)
	if expirationTime.Unix() < actionTime.Unix() {
		if lastaction.Add(expDuration).Unix() > actionTime.Unix() {
			claims["created"] = actionTime
			at := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
			newToken, err := at.SignedString(([]byte)(os.Getenv("TOKEN_SECRET")))
			if err != nil {
				s.UpdateUserLastAction(int32(user), actionTime)
				return credential, "", true
			}
			s.UpdateUserLastAction(int32(user), actionTime)
			return credential, newToken, true
		}
	} else {
		s.UpdateUserLastAction(int32(user), actionTime)
		return credential, "", true
	}
	return nil, "", true
}

// Validations

// ValidateTeacherActive validate teacher credentials
func (s *Service) ValidateTeacherActive(c *UserCredential) bool {
	if c.Rol != 2 {
		return false
	}
	if c.Status != "active" {
		return false
	}
	return true
}
