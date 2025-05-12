package repoadapter

import (
	"github.com/alfariiizi/go-service/internal/domain/entity"
	"github.com/alfariiizi/go-service/internal/domain/model"
	"github.com/alfariiizi/go-service/internal/infrastructure/database"
	"github.com/alfariiizi/go-service/internal/repository/repoport"
)

type userRepository struct {
	db *database.GormDB
}

func NewUserRepository(db *database.GormDB) repoport.UserRepository {
	db.DB.AutoMigrate(&entity.User{})
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers() ([]entity.User, error) {
	// Implementation for getting all users from the database
	var users []entity.User
	if err := r.db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id uint) (entity.User, error) {
	// Implementation for getting a user by ID from the database
	var user entity.User
	if err := r.db.DB.First(&user, id).Error; err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user model.UserRequest) (entity.User, error) {
	// Implementation for creating a new user in the database
	if err := r.db.DB.Create(&user).Error; err != nil {
		return entity.User{}, err
	}
	var createdUser entity.User
	r.db.DB.Last(&createdUser)
	return createdUser, nil
}

func (r *userRepository) UpdateUser(id uint, user model.UserRequest) (entity.User, error) {
	// Implementation for updating an existing user in the database
	if err := r.db.DB.Save(&user).Error; err != nil {
		return entity.User{}, err
	}
	var updatedUser entity.User
	r.db.DB.First(&updatedUser, id)
	return updatedUser, nil
}

func (r *userRepository) DeleteUser(id uint) error {
	// Implementation for deleting a user by ID from the database
	if err := r.db.DB.Delete(&entity.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
