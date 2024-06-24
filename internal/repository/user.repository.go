package repository

import (
	"gorm.io/gorm"

	"github.com/WasathTheekshana/golang-project-structure/interfaces"
	"github.com/WasathTheekshana/golang-project-structure/internal/model"
)

type UserRepository interface {
	CreateUserAccount(userRequest *interfaces.IUserRegistrationRequest) (*model.User, error)
	FetchUserDetails(userEmail string) (*model.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

// Function to create a new user
func (r *userRepository) CreateUserAccount(userRequest *interfaces.IUserRegistrationRequest) (*model.User, error) {
	user := &model.User{
		Email:    userRequest.Email,
		Username: userRequest.Username,
		FullName: userRequest.FullName,
		UserRole: model.UserRole,
	}

	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

// Method to return user credentials
func (r *userRepository) FetchUserDetails(userEmail string) (*model.User, error) {
	user := &model.User{}

	if err := r.db.Where("email = ?", userEmail).First(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
