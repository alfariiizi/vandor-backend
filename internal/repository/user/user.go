package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/alfariiizi/go-service/database/db"
	"github.com/alfariiizi/go-service/internal/model"
	"github.com/google/uuid"
)

type userRepository struct {
	db *db.Queries
}

func NewUserRepository(db *db.Queries) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers() ([]model.UserResponse, error) {
	users, err := r.db.GetAllUsers(context.Background())
	if err != nil {
		return nil, err
	}
	return model.ToUserResponseList(users), nil
}

func (r *userRepository) GetUserByID(id string) (model.UserResponse, error) {
	parseId, err := uuid.Parse(id)
	if err != nil {
		return model.UserResponse{}, fmt.Errorf("invalid user id: %w", err)
	}
	user, err := r.db.GetUserById(context.Background(), parseId)
	if err != nil {
		return model.UserResponse{}, fmt.Errorf("failed to get user by id: %w", err)
	}

	return model.ToUserResponse(&user), nil
}

func (r *userRepository) CreateUser(input model.UserRequest) (model.UserResponse, error) {
	id := uuid.New()

	user, err := r.db.CreateUser(context.Background(), db.CreateUserParams{
		ID:       id,
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return model.ToUserResponse(nil), err
	}

	return model.ToUserResponse(&user), nil
}

func (r *userRepository) UpdateUser(id string, input model.UserRequest) (model.UserResponse, error) {
	parseId, err := uuid.Parse(id)
	if err != nil {
		return model.UserResponse{}, fmt.Errorf("invalid user id: %w", err)
	}

	user, err := r.db.UpdateUser(context.Background(), db.UpdateUserParams{
		ID:       parseId,
		Username: sql.NullString{String: input.Username, Valid: true},
		Email:    sql.NullString{String: input.Email, Valid: true},
		Password: sql.NullString{String: input.Password, Valid: true},
	})
	if err != nil {
		return model.ToUserResponse(nil), err
	}

	return model.ToUserResponse(&user), nil
}

func (r *userRepository) DeleteUser(id string) error {
	parseId, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid user id: %w", err)
	}
	_, err = r.db.DeleteUser(context.Background(), parseId)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
