package user

import (
	"github.com/alfariiizi/go-service/internal/domain/entity"
	"github.com/alfariiizi/go-service/internal/domain/model"
	"github.com/alfariiizi/go-service/internal/repository/user"
)

type userServiceAdapter struct {
	userRepo user.UserRepository
}

func NewUserService(userRepo user.UserRepository) UserService {
	return &userServiceAdapter{
		userRepo: userRepo,
	}
}

func (u *userServiceAdapter) ListUsers() ([]entity.User, error) {
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userServiceAdapter) CreateUser(user model.UserRequest) (entity.User, error) {
	res, err := u.userRepo.CreateUser(user)
	return res, err
}

func (u *userServiceAdapter) GetUser(id uint) (entity.User, error) {
	user, err := u.userRepo.GetUserByID(id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (u *userServiceAdapter) UpdateUser(id uint, user model.UserRequest) (entity.User, error) {
	updatedUser, err := u.userRepo.UpdateUser(id, user)
	if err != nil {
		return entity.User{}, err
	}
	return updatedUser, nil
}

func (u *userServiceAdapter) DeleteUser(id uint) error {
	err := u.userRepo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
