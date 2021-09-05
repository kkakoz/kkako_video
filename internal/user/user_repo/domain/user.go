package domain

import (
	"context"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Brief    string `json:"brief"`
	Salt     string `json:"salt"`
}

type IUserRepo interface {
	AddUser(ctx context.Context, user *User) error
	GetUserById(ctx context.Context, id int64) (*User, error)
	GetUserList(ctx context.Context, ids []int64) ([]*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
