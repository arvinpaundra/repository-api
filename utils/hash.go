package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	defaultCost := bcrypt.DefaultCost

	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), defaultCost)

	return string(hashed)
}

func ComparePassword(hashedPassword string, password string) bool {
	valid := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return valid != nil
}
