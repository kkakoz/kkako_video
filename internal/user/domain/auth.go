package domain

import (
	"context"
)

type Auth struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Salt     string `json:"salt"`
	UserId   int64  `json:"user_id"`
	Auth     int32  `json:"auth"`
}

type IAuthRepo interface {
	AddAuth(ctx context.Context, auth *Auth) error
	GetAuth(ctx context.Context, id int64) (*Auth, error)
	DeleteAuth(ctx context.Context, id int64) error
	GetAuthByEmail(ctx context.Context, email string) (*Auth, error)
}
