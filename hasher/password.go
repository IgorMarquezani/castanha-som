package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(passwd string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwd), 10)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func CompareHashAndPassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
