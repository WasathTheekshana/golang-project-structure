package interfaces

import (
	"time"

	"github.com/google/uuid"

	"github.com/WasathTheekshana/golang-project-structure/internal/model"
)

type IUserRegistrationRequest struct {
	Email    string `json:"email" validate:"required,email"`
	FullName string `json:"fullName" validate:"required"`
	Username string `json:"username" validate:"required"`
}

type IUserData struct {
	ID        uuid.UUID  `json:"id"`
	Email     string     `json:"email"`
	Username  string     `json:"username"`
	FullName  string     `json:"fullName"`
	UserRole  model.Role `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
}
