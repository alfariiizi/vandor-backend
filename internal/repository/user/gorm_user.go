package user

import (
	"fmt"
	"time"

	"github.com/alfariiizi/go-service/internal/domain/entity"
	"github.com/alfariiizi/go-service/internal/domain/model"
	"github.com/alfariiizi/go-service/internal/infrastructure/database"
	"github.com/google/uuid"
)

type userRepository struct {
	db *database.GormDB
}

func NewUserRepository(db *database.GormDB) UserRepository {
	db.DB.Migrator().DropTable(&entity.User{}) // only in dev!
	db.DB.AutoMigrate(&entity.User{})
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers() ([]model.UserResponse, error) {
	// Implementation for getting all users from the database
	var users []entity.User
	if err := r.db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return model.ToUserResponseList(users), nil
}

func (r *userRepository) GetUserByID(id uint) (model.UserResponse, error) {
	// Implementation for getting a user by ID from the database
	var user entity.User
	if err := r.db.DB.First(&user, id).Error; err != nil {
		return model.ToUserResponse(nil), err
	}
	return model.ToUserResponse(&user), nil
}

func (r *userRepository) CreateUser(input model.UserRequest) (model.UserResponse, error) {
	// Check if the user already exists
	var existingUser entity.User
	if err := r.db.DB.Where("username = ? OR email = ?", input.Username, input.Email).First(&existingUser).Error; err == nil {
		return model.ToUserResponse(nil), fmt.Errorf("user with username or email already exists")
	}

	id := uuid.New().String()
	now := time.Now()

	if err := r.db.DB.Create(entity.User{
		ID:        id,
		Username:  input.Username,
		Email:     input.Email,
		Password:  input.Password,
		CreatedAt: now,
		UpdatedAt: now,
	}).Error; err != nil {
		return model.ToUserResponse(nil), err
	}
	var createdUser entity.User
	r.db.DB.Last(&createdUser)
	return model.ToUserResponse(&createdUser), nil
}

func (r *userRepository) UpdateUser(id uint, user model.UserRequest) (model.UserResponse, error) {
	// Implementation for updating an existing user in the database
	if err := r.db.DB.Save(&user).Error; err != nil {
		return model.ToUserResponse(nil), err
	}
	var updatedUser entity.User
	r.db.DB.First(&updatedUser, id)
	return model.ToUserResponse(&updatedUser), nil
}

func (r *userRepository) DeleteUser(id uint) error {
	// Implementation for deleting a user by ID from the database
	if err := r.db.DB.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
