package services

import (
	"github.com/WasathTheekshana/golang-project-structure/interfaces"
	"github.com/WasathTheekshana/golang-project-structure/internal/repository"
)

type UserService interface {
	FetchUserAccount(userEmail string) (*interfaces.IUserData, error)
	CreateUserAccount(userRequest *interfaces.IUserRegistrationRequest) (*interfaces.IUserData, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

// Method to create a user
func (s *userService) CreateUserAccount(userRequest *interfaces.IUserRegistrationRequest) (*interfaces.IUserData, error) {
	userData, err := s.userRepo.CreateUserAccount(userRequest)
	if err != nil {
		return nil, err
	}

	return &interfaces.IUserData{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		Username:  userData.Username,
		UserRole:  userData.UserRole,
		CreatedAt: userData.CreatedAt,
	}, nil
}

// Fetch user details
func (s *userService) FetchUserAccount(userEmail string) (*interfaces.IUserData, error) {
	userData, err := s.userRepo.FetchUserDetails(userEmail)
	if err != nil {
		return nil, err
	}

	return &interfaces.IUserData{
		ID:        userData.ID,
		FullName:  userData.FullName,
		Email:     userData.Email,
		Username:  userData.Username,
		UserRole:  userData.UserRole,
		CreatedAt: userData.CreatedAt,
	}, nil
}
