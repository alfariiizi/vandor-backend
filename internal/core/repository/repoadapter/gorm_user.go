package repoadapter

import (
	"github.com/alfariiizi/go-service/internal/core/domain"
	"github.com/alfariiizi/go-service/internal/core/repository/repoport"
	"github.com/alfariiizi/go-service/internal/infrastructure/database"
)

type userRepository struct {
	db *database.GormDB
}

func NewUserRepository(db *database.GormDB) repoport.UserRepository {
	db.DB.AutoMigrate(&domain.User{})
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAllUsers() ([]domain.User, error) {
	// Implementation for getting all users from the database
	var users []domain.User
	if err := r.db.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetUserByID(id uint) (domain.User, error) {
	// Implementation for getting a user by ID from the database
	var user domain.User
	if err := r.db.DB.First(&user, id).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *userRepository) CreateUser(user domain.UserRequest) (domain.User, error) {
	// Implementation for creating a new user in the database
	if err := r.db.DB.Create(&user).Error; err != nil {
		return domain.User{}, err
	}
	var createdUser domain.User
	r.db.DB.Last(&createdUser)
	return createdUser, nil
}

func (r *userRepository) UpdateUser(id uint, user domain.UserRequest) (domain.User, error) {
	// Implementation for updating an existing user in the database
	if err := r.db.DB.Save(&user).Error; err != nil {
		return domain.User{}, err
	}
	var updatedUser domain.User
	r.db.DB.First(&updatedUser, id)
	return updatedUser, nil
}

func (r *userRepository) DeleteUser(id uint) error {
	// Implementation for deleting a user by ID from the database
	if err := r.db.DB.Delete(&domain.User{}, id).Error; err != nil {
		return err
	}
	return nil
}
