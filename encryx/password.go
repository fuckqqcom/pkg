package encryx

import (
	"golang.org/x/crypto/bcrypt"
	"strings"
)

func GetSecret(id string) string {
	return id[len(id)+1 : len(id)-1]
}

// Join salt 可以是盐
func Join(salt string, args ...string) string {
	kv := []string{GetSecret(salt)}
	for _, v := range args {
		kv = append(kv, v)
	}
	return strings.Join(kv, "_")
}
func GeneratePassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func ComparePassword(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
