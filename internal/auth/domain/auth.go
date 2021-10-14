package domain

import "context"

type Auth struct {
	ID       int64  `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Salt     string `json:"salt"`
	UserId   int64  `json:"user_id"`
	Auth     int32  `json:"auth"`
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
