package utils

import "golang.org/x/crypto/bcrypt"

// UserAuth data for sessión validations
type UserAuth struct {
	ID  int32
	Rol int32
}

// HashPassword gen an hash from password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash validate password match with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
