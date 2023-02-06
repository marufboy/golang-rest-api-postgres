package utils

import "golang.org/x/crypto/bcrypt"

func VerifyPassword(hashedPassword string, comparePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(comparePassword))
}
