package user

import (
	"github.com/alfariiizi/go-service/internal/domain/model"
)

type UserRepository interface {
	// GetAllUsers retrieves all users from the repository.
	GetAllUsers() ([]model.UserResponse, error)
	// GetUserByID retrieves a user by their ID from the repository.
	GetUserByID(id string) (model.UserResponse, error)
	// CreateUser creates a new user in the repository.
	CreateUser(user model.UserRequest) (model.UserResponse, error)
	// UpdateUser updates an existing user in the repository.
	UpdateUser(id string, user model.UserRequest) (model.UserResponse, error)
	// DeleteUser deletes a user by their ID from the repository.
	DeleteUser(id string) error
}
