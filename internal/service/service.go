package service

import (
	"github.com/alfariiizi/go-service/internal/repository"
	"github.com/alfariiizi/go-service/internal/service/user"
)

type Services struct {
	User user.UserService
}

func InitServices(
	repos *repository.Repositories,
) Services {
	return Services{
		User: user.NewUserService(repos.User),
	}
}
