package handler

import (
	"context"
	v1 "kkako_video/api/user/v1"
	"kkako_video/internal/user/domain"
	"kkako_video/pkg/db/mysqlx"
)

type UserHandler struct {
	v1.UnimplementedUserServiceServer
	v1.UnimplementedLoginServiceServer
}

func (u UserHandler) Register(ctx context.Context, req *v1.RegisterReq) (*v1.RegisterRes, error) {
	panic("implement me")
}

func (u UserHandler) Login(ctx context.Context, req *v1.LoginReq) (*v1.LoginRes, error) {
	panic("implement me")
}

func (u UserHandler) UserInfo(ctx context.Context, req *v1.GetUserReq) (*v1.GetUserRes, error) {
	panic("implement me")
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

func NewUserRepoHandler() *UserHandler {
	db := mysqlx.GetDB(context.TODO())
	db.AutoMigrate(&domain.User{})
	return &UserHandler{}
}
