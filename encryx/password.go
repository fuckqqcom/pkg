package encryx

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func GetSecret(id string) string {
	return id[6:16]
}

func Password(password, id string) string {
	return strings.Join([]string{password, GetSecret(id)}, "_")
}
func GeneratePassword(password, id string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Password(password, id)), 10)
	return string(bytes), err
}

func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
