package handler

import (
	"context"
	"github.com/jinzhu/copier"
	v1 "kkako_video/api/user/v1"
	"kkako_video/internal/user/domain"
)

var _ v1.UserServiceServer = (*UserHandler)(nil)

func NewUserHandler(userLogic domain.IUserLogic) *UserHandler {
	return &UserHandler{userLogic: userLogic}
}

type UserHandler struct {
	v1.UnimplementedUserServiceServer
	userLogic domain.IUserLogic
}

func (u UserHandler) GetUser(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserRes, error) {
	user, err := u.userLogic.GetUser(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	res := &v1.GetUserRes{}
	err = copier.Copy(res, user)
	return res, err
}

func (u UserHandler) UserNews(ctx context.Context, req *v1.UserNewsReq) (*v1.UserNewsRes, error) {
	panic("implement me")
}

func (u UserHandler) UserFollow(ctx context.Context, req *v1.UserFollowReq) (*v1.UserFollowRes, error) {
	panic("implement me")
}

func (u UserHandler) UserLike(ctx context.Context, req *v1.UserLikeReq) (*v1.UserLikeRes, error) {
	panic("implement me")
}

func (u UserHandler) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterRes, error) {
	auth := &domain.Auth{
		Email:    req.Email,
		Password: req.Password,
		Name:     req.Name,
	}
	err := u.userLogic.Register(ctx, auth)
	return &v1.RegisterRes{}, err
}

func (u UserHandler) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginRes, error) {
	user := &domain.Auth{
		Email:    req.Email,
		Password: req.Password,
	}
	token, err := u.userLogic.Login(ctx, user)
	return &v1.LoginRes{Token: token}, err
}
