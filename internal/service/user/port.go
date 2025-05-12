package user

import (
	"github.com/alfariiizi/go-service/internal/domain/model"
)

type UserService interface {
	ListUsers() ([]model.UserResponse, error)
	GetUser(id string) (model.UserResponse, error)
	CreateUser(user model.UserRequest) (model.UserResponse, error)
	UpdateUser(id string, user model.UserRequest) (model.UserResponse, error)
	DeleteUser(id string) error
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
