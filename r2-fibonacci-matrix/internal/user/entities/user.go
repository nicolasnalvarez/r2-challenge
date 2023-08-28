package entities

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User struct is used to store user information in the database
type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

// HashPassword encrypts user password using bcrypt
// It returns an error if there is an issue encrypting the password
func (u *User) HashPassword() error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	u.Password = string(bytes)
	return err
}

// CheckPassword checks user password
// CheckPassword takes a string as a parameter and compares it to the user's encrypted password
// It returns an error if there is an issue comparing the passwords
func (u *User) CheckPassword(providedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(providedPassword))
}
