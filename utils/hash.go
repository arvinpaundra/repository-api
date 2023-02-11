package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashed)
}

func ComparePassword(hashedPassword string, password string) bool {
	ok := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return ok == nil
}
