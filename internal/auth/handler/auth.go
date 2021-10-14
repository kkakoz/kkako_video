package handler

import (
	"context"
	v1 "kkako_video/api/user/v1"
	"kkako_video/internal/auth/domain"
)

type AuthHandler struct {
	v1.UnimplementedAuthServiceServer
	userLogic domain.IAuthLogic
}


func (u AuthHandler) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterRes, error) {
	user := &domain.Auth{
		Email: req.Email,
		Password: req.Password,
	}
	err := u.userLogic.Register(ctx, user, req.Name)
	return &v1.RegisterRes{}, err
}

func (u AuthHandler) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginRes, error) {
	user := &domain.Auth{
		Email: req.Email,
		Password: req.Password,
	}
	_, _, err := u.userLogic.Login(ctx, user)
	return &v1.LoginRes{}, err
}

func NewAuthHandler(userLogic domain.IAuthLogic) *AuthHandler {
	return &AuthHandler{userLogic: userLogic}
}
