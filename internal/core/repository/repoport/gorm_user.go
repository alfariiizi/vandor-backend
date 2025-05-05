package repoport

import "github.com/alfariiizi/go-service/internal/core/domain"

type UserRepository interface {
	// GetAllUsers retrieves all users from the repository.
	GetAllUsers() ([]domain.User, error)
	// GetUserByID retrieves a user by their ID from the repository.
	GetUserByID(id uint) (domain.User, error)
	// CreateUser creates a new user in the repository.
	CreateUser(user domain.UserRequest) (domain.User, error)
	// UpdateUser updates an existing user in the repository.
	UpdateUser(id uint, user domain.UserRequest) (domain.User, error)
	// DeleteUser deletes a user by their ID from the repository.
	DeleteUser(id uint) error
}
