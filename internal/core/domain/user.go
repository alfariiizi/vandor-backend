package domain

type User struct {
	ID        uint
	Username  string
	Email     *string
	Password  string
	Role      string
	CreatedAt string
	UpdatedAt string
}

type UserRequest struct {
	Username string  `json:"username" validate:"required"`
	Email    *string `json:"email" validate:"required,email"`
	Password string  `json:"password" validate:"required"`
	Role     string  `json:"role" validate:"required"`
}

type UserResponse struct {
	ID        uint    `json:"id"`
	Username  string  `json:"username"`
	Email     *string `json:"email"`
	Role      string  `json:"role"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

func ToResponse(u User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Username:  u.Username,
		Email:     u.Email,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
