package model

import (
	"time"

	"github.com/alfariiizi/go-service/internal/domain/entity"
)

type UserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func ToUserResponse(u *entity.User) UserResponse {
	if u == nil {
		return UserResponse{}
	}

	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToUserResponseList(users []entity.User) []UserResponse {
	userResponses := make([]UserResponse, len(users))
	for i, u := range users {
		userResponses[i] = ToUserResponse(&u)
	}
	return userResponses
}
