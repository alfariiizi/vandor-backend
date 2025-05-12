package user

import (
	"github.com/alfariiizi/go-service/internal/domain/entity"
	"github.com/alfariiizi/go-service/internal/domain/model"
)

type UserRepository interface {
	// GetAllUsers retrieves all users from the repository.
	GetAllUsers() ([]entity.User, error)
	// GetUserByID retrieves a user by their ID from the repository.
	GetUserByID(id uint) (entity.User, error)
	// CreateUser creates a new user in the repository.
	CreateUser(user model.UserRequest) (entity.User, error)
	// UpdateUser updates an existing user in the repository.
	UpdateUser(id uint, user model.UserRequest) (entity.User, error)
	// DeleteUser deletes a user by their ID from the repository.
	DeleteUser(id uint) error
}
