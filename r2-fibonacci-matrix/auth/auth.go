package auth

import (
	"errors"
	jwt "github.com/golang-jwt/jwt/v5"
	"time"
)

// JwtWrapper wraps the signing key and the issuer
// JwtWrapper is a struct that holds the secret key, issuer and expiration time for a JWT token
type JwtWrapper struct {
	SecretKey         string // key used for signing the JWT token
	Issuer            string // Issuer of the JWT token
	ExpirationMinutes int64  // Number of minutes the JWT token will be valid for
	ExpirationHours   int64  // Expiration time of the JWT token in hours
}

// JwtClaim adds email as a claim to the token
// JwtClaim is a struct that holds the Email of the user, as well as the StandardClaims
type JwtClaim struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken generates a jwt token
// GenerateToken takes an email as an argument and returns a signed JWT token and an error
func (j *JwtWrapper) GenerateToken(email string) (string, error) {
	claims := &JwtClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(j.ExpirationMinutes))),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SecretKey))

}

// RefreshToken generates a refresh jwt token
// RefreshToken takes an email as an argument and returns a signed JWT token and an error
func (j *JwtWrapper) RefreshToken(email string) (string, error) {
	claims := &JwtClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Hour * time.Duration(j.ExpirationHours))),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.SecretKey))
}

// ValidateToken validates the jwt token
// ValidateToken takes a signed JWT token as an argument and returns the JwtClaim and an error
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}
	if claims.ExpiresAt.Before(time.Now().Local()) {
		err = errors.New("JWT is expired")
		return
	}
	return
}
