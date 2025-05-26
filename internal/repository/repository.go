package repository

import (
	"github.com/alfariiizi/go-service/database/db"
	"github.com/alfariiizi/go-service/internal/repository/user"
)

type Repositories struct {
	User user.UserRepository
}

func InitRepositories(
	db *db.Queries,
) *Repositories {
	return &Repositories{
		User: user.NewUserRepository(db),
	}
}
