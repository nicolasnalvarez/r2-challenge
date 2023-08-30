package auth

import (
	"errors"
	jwt "github.com/golang-jwt/jwt/v5"
	"time"
)

// Service wraps the signing key and the issuer
// Service is a struct that holds the secret key, issuer and expiration time for a JWT token
type (
	Service struct {
		accessTokenSecretKey          string // key used for signing the JWT token
		refreshTokenSecretKey         string // key used for signing the refresh token
		issuer                        string // Issuer of the JWT token
		accessTokenExpirationMinutes  int    // Number of minutes the JWT token will be valid for
		refreshTokenExpirationMinutes int    // Expiration time of the JWT token in hours
	}

	JwtService interface {
		GenerateToken(email string) (string, error)
		RefreshToken(email string) (string, error)
		ValidateAccessToken(signedToken string) (claims *JwtClaim, err error)
		ValidateRefreshToken(signedToken string) (claims *JwtClaim, err error)
	}

	JwtClaim struct {
		Email string `json:"email"`
		jwt.RegisteredClaims
	}
)

func NewJwtService(accessTokenSecretKey, refreshTokenSecretKey, issuer string,
	accessTokenExpirationMinutes, refreshTokenExpirationMinutes int) *Service {
	return &Service{
		accessTokenSecretKey:          accessTokenSecretKey,
		refreshTokenSecretKey:         refreshTokenSecretKey,
		issuer:                        issuer,
		accessTokenExpirationMinutes:  accessTokenExpirationMinutes,
		refreshTokenExpirationMinutes: refreshTokenExpirationMinutes,
	}
}

// GenerateToken generates a jwt token
// GenerateToken takes an email as an argument and returns a signed JWT token and an error
func (s *Service) GenerateToken(email string) (string, error) {
	claims := &JwtClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(s.accessTokenExpirationMinutes))),
			Issuer:    s.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.accessTokenSecretKey))

}

// RefreshToken generates a refresh jwt token
// RefreshToken takes an email as an argument and returns a signed JWT token and an error
func (s *Service) RefreshToken(email string) (string, error) {
	claims := &JwtClaim{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(s.refreshTokenExpirationMinutes))),
			Issuer:    s.issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.refreshTokenSecretKey))
}

// ValidateAccessToken validates the jwt token
// ValidateAccessToken takes a signed JWT token as an argument and returns the JwtClaim and an error
func (s *Service) ValidateAccessToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.accessTokenSecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt.Before(time.Now().Local()) {
		err = errors.New("JWT is expired")
		return
	}
	return
}

// ValidateRefreshToken validates the jwt refresh token
// ValidateRefreshToken takes a signed JWT refresh token as an argument and returns the JwtClaim and an error
func (s *Service) ValidateRefreshToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(s.refreshTokenSecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt.Before(time.Now().Local()) {
		err = errors.New("JWT is expired")
		return
	}
	return
}
