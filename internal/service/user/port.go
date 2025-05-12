package user

import (
	"github.com/alfariiizi/go-service/internal/domain/entity"
	"github.com/alfariiizi/go-service/internal/domain/model"
)

type UserService interface {
	// ListUsers retrieves a list of all users.
	ListUsers() ([]entity.User, error)
	// GetUser retrieves a user by their ID.
	GetUser(id uint) (entity.User, error)
	// CreateUser creates a new user with the given details.
	CreateUser(user model.UserRequest) (entity.User, error)
	// UpdateUser updates the details of an existing user.
	UpdateUser(id uint, user model.UserRequest) (entity.User, error)
	// DeleteUser deletes a user by their ID.
	DeleteUser(id uint) error
	// // AuthenticateUser authenticates a user with the given credentials.
	// AuthenticateUser(username, password string) (domain.UserResponse, error)
	// // ChangePassword changes the password of a user.
	// ChangePassword(id, oldPassword, newPassword string) error
	// // ResetPassword resets the password of a user.
	// ResetPassword(id, newPassword string) error
	// // GetUserByEmail retrieves a user by their email address.
	// GetUserByEmail(email string) (domain.UserResponse, error)
	// // GetUserByUsername retrieves a user by their username.
	// GetUserByUsername(username string) (domain.UserResponse, error)
	// // GetUserByRole retrieves a user by their role.
	// GetUserByRole(role string) ([]domain.UserResponse, error)
}
