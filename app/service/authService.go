package service

import (
	"github.com/vutuankiet4599/go-jwt/app/models"
	"github.com/vutuankiet4599/go-jwt/app/repository"
	"github.com/vutuankiet4599/go-jwt/app/request"
	"github.com/vutuankiet4599/go-jwt/helper"
)

type AuthService interface{
	IsDuplicateEmail(email string) bool
	CreateUser(registerRequest request.RegisterRequest) (*models.User, bool, string)
	VerifyCredentials(loginRequest request.LoginRequest) (*models.User, bool, string)
	GetCurrentUser(id uint) (*models.User, bool, string)
}

type authService struct {
	userRepository repository.UserRepository
}

func (s *authService) CreateUser(registerRequest request.RegisterRequest) (*models.User, bool, string) {
	response, user := s.userRepository.CreateUser(registerRequest.Email, registerRequest.Password, registerRequest.Name)
	if response.Error != nil {
		return &models.User{}, true, response.Error.Error()
	}
	return user, false, ""
}

func (s *authService) IsDuplicateEmail(email string) bool {
	res := s.userRepository.IsDuplicateEmail(email)
	return (res != nil)
}

func (s *authService) VerifyCredentials(loginRequest request.LoginRequest) (*models.User, bool, string) {
	response, user := s.userRepository.VerifyCredentials(loginRequest.Email)
	if response.Error != nil {
		return &models.User{}, true, response.Error.Error()
	}
	if !helper.CompareHashValue(user.Password, loginRequest.Password) {
		return &models.User{}, true, "Password is incorrect"

	}
	return user, false, ""
}

func (s *authService) GetCurrentUser(id uint) (*models.User, bool, string) {
	response, user := s.userRepository.GetUser(id)
	if response.Error != nil {
		return &models.User{}, true, response.Error.Error()
	}
	return user, false, ""
}

func NewAuthService(userRepository repository.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}
