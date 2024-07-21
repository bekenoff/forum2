package pkg

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(psw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(psw), bcrypt.DefaultCost)
	return string(bytes), err
}
