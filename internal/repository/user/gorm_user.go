package user

import (
	"fmt"
	"time"

	"github.com/alfariiizi/go-service/internal/domain/entity"
	"github.com/alfariiizi/go-service/internal/domain/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	// db.Migrator().DropTable(&entity.User{}) // only in dev!
	db.AutoMigrate(&entity.User{})
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers() ([]model.UserResponse, error) {
	var users []entity.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return model.ToUserResponseList(users), nil
}

func (r *userRepository) GetUserByID(id string) (model.UserResponse, error) {
	var user entity.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		return model.ToUserResponse(nil), err
	}
	return model.ToUserResponse(&user), nil
}

func (r *userRepository) CreateUser(input model.UserRequest) (model.UserResponse, error) {
	var existingUser entity.User
	if err := r.db.Where("username = ? OR email = ?", input.Username, input.Email).First(&existingUser).Error; err == nil {
		return model.ToUserResponse(nil), fmt.Errorf("user with username or email already exists")
	}

	id := uuid.New().String()
	now := time.Now()

	if err := r.db.Create(entity.User{
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
	r.db.Last(&createdUser)
	return model.ToUserResponse(&createdUser), nil
}

func (r *userRepository) UpdateUser(id string, user model.UserRequest) (model.UserResponse, error) {
	if err := r.db.Save(&user).Error; err != nil {
		return model.ToUserResponse(nil), err
	}
	var updatedUser entity.User
	r.db.First(&updatedUser, id)
	return model.ToUserResponse(&updatedUser), nil
}

func (r *userRepository) DeleteUser(id string) error {
	if err := r.db.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
