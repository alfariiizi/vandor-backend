package domain

import (
	"fmt"

	domain_builder "github.com/alfariiizi/vandor/internal/core/domain/builder"
	"github.com/alfariiizi/vandor/internal/infrastructure/db"
)

type User struct {
	*db.User
	client *db.Client
}

func NewUserDomain(client *db.Client) domain_builder.Domain[*db.User, *User] {
	return domain_builder.NewDomain(
		func(e *db.User, c *db.Client) *User {
			return &User{
				User:   e,
				client: c,
			}
		}, client)
}

func (u *User) CanLogin() bool {
	return u.Email != ""
}

func (u *User) FullName() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}
