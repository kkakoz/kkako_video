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
	Follows  []int64 `json:"follows"`
	Fans     []int64 `json:"fans"`
	Likes    []int64 `json:"likes"`
	State    int32   `json:"state"`
}

type IUserLogic interface {
	AddUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUsers(ctx context.Context, ids []int64) ([]*User, error)
}

type IUserRepo interface {
	AddUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUserList(ctx context.Context, ids []int64) ([]*User, error)
	GetUserByCond(ctx context.Context, email string) (*User, error)
}
