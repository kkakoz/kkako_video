package domain

import (
	"context"
)

type User struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Avatar      string `json:"avatar"`
	Brief       string `json:"brief"`
	FollowCount int64  `json:"follow_count"`
	FansCount   int64  `json:"fans_count"`
	LikeCount   int64  `json:"like_count"`
	State       int32  `json:"state"`
	Auth        *Auth  `json:"auth"`
}

type IUserLogic interface {
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUsers(ctx context.Context, ids []int64) ([]*User, error)
	Register(ctx context.Context, auth *Auth) error
	Login(ctx context.Context, user *Auth) (string, error)
}

type IUserRepo interface {
	AddUser(ctx context.Context, user *User) error
	GetUser(ctx context.Context, id int64) (*User, error)
	GetUserList(ctx context.Context, ids []int64) ([]*User, error)
}
