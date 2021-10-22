package domain

import (
	"context"
)

type Auth struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Salt     string `json:"salt"`
	UserId   int64  `json:"user_id"`
	Auth     int32  `json:"auth"`
}

type User struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Avatar      string  `json:"avatar"`
	Brief       string  `json:"brief"`
	FollowCount int64   `json:"follow_count"`
	FansCount   int64   `json:"fans_count"`
	LikeCount   int64   `json:"like_count"`
	State       int32   `json:"state"`
}

type IAuthLogic interface {
	Register(ctx context.Context, user *Auth, name string) error
	Login(ctx context.Context, user *Auth) (int64, string, error)
}

type IAuthRepo interface {
	AddAuth(ctx context.Context, auth *Auth) error
	GetAuth(ctx context.Context, id int64) (*Auth, error)
	DeleteAuth(ctx context.Context, id int64) error
	GetAuthByEmail(ctx context.Context, email string) (*Auth, error)
}
