package auth

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword encrypts user password
// HashPassword takes a string as a parameter and encrypts it using bcrypt
// It returns an error if there is an issue encrypting the password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPassword checks user password
// CheckPassword takes a string as a parameter and compares it to the user's encrypted password
// It returns an error if there is an issue comparing the passwords
func CheckPassword(providedPassword, userPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(providedPassword))
}
