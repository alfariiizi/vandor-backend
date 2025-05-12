package entity

type User struct {
	ID        uint
	Username  string
	Email     *string
	Password  string
	Role      string
	CreatedAt string
	UpdatedAt string
}
