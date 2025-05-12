package repository

import (
	"github.com/alfariiizi/go-service/internal/infrastructure/database"
	"github.com/alfariiizi/go-service/internal/repository/user"
)

type Repositories struct {
	User user.UserRepository
}

func InitRepositories(
	db *database.GormDB,
) *Repositories {
	return &Repositories{
		User: user.NewUserRepository(db),
	}
}
