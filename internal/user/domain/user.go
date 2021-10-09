package domain

import "context"

type User struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Password string  `json:"password"`
	Email    string  `json:"email"`
	Avatar   string  `json:"avatar"`
	Brief    string  `json:"brief"`
	Salt     string  `json:"salt"`
}

type IUserService interface {
	Register(ctx context.Context, user *User) error
	Login(ctx context.Context, user *User) (string, error)
	Get(ctx context.Context, id int64) (*User, error)
	GetList(ctx context.Context, ids []int64) ([]*User, error)
}

type IUserRepo interface {
	Get(ctx context.Context, id int64) (*User, error)
	GetList(ctx context.Context, ids []int64) ([]*User, error)
}