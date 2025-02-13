package hash

import "golang.org/x/crypto/bcrypt"

func PasswordHash(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed), err
}
