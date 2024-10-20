package types

import "time"

type UserStore interface {
	GetUserByUsername(username string) (*User, error)
	GetUserByID(id int) (*User, error)
	CreateUser(*User) error
}

// User Entity
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}

// 注册用户请求体
type RegisterUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}
