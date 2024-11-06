package bcrypt

import (
	"golang.org/x/crypto/bcrypt"
)

type BCryptHasher struct {
}

func New() *BCryptHasher {
	return &BCryptHasher{}
}

func (hsh BCryptHasher) HashPassword(password string, cost int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
func (hsh BCryptHasher) ComparePasswordAndHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
