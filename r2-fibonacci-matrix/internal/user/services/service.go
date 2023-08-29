package services

import (
	"r2-fibonacci-matrix/auth"
	"r2-fibonacci-matrix/internal/user/dtos"
	"r2-fibonacci-matrix/internal/user/entities"
	"r2-fibonacci-matrix/internal/user/repositories"
)

type (
	Service struct {
		userRepository repositories.UserRepository
	}

	UserService interface {
		Save(registerRequest dtos.RegisterRequest) error
		Login(loginRequest dtos.LoginRequest) (dtos.LoginResponse, error)
		FindUserByEmail(email string) (entities.User, error)
	}
)

func NewUserService(userRepository repositories.UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) Save(registerRequest dtos.RegisterRequest) error {
	user := entities.User{
		Name:     registerRequest.Name,
		Email:    registerRequest.Email,
		Password: registerRequest.Password,
	}
	if err := user.HashPassword(); err != nil {
		return err
	}
	return s.userRepository.SaveUser(user)
}

func (s *Service) Login(loginRequest dtos.LoginRequest) (dtos.LoginResponse, error) {
	user, err := s.FindUserByEmail(loginRequest.Email)
	if err != nil {
		return dtos.LoginResponse{}, err
	}
	err = user.CheckPassword(loginRequest.Password)
	if err != nil {
		return dtos.LoginResponse{}, err
	}

	// values used just to test the app freely
	jwtWrapper := auth.JwtWrapper{
		SecretKey:         "verysecretkey",
		Issuer:            "r2",
		ExpirationMinutes: 60,
		ExpirationHours:   12,
	}
	token, err := jwtWrapper.GenerateToken(user.Email)
	if err != nil {
		return dtos.LoginResponse{}, err
	}
	refreshToken, err := jwtWrapper.RefreshToken(user.Email)
	if err != nil {
		return dtos.LoginResponse{}, err
	}

	return dtos.LoginResponse{
		Token:        token,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) FindUserByEmail(email string) (entities.User, error) {
	return s.userRepository.FindUserByEmail(email)
}
