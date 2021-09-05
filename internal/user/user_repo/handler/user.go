package handler

import (
	"context"
	"github.com/jinzhu/copier"
	v1 "kkako_video/api/user/v1"
	"kkako_video/internal/user/user_repo/domain"
)

type UserRepoHandler struct {
	v1.UnimplementedUserRepoServiceServer
	userRepo domain.IUserRepo
}

func NewUserRepoHandler(userRepo domain.IUserRepo) *UserRepoHandler {
	return &UserRepoHandler{userRepo: userRepo}
}

func (u UserRepoHandler) AddUser(ctx context.Context, req *v1.AddUserReq) (*v1.AddUserRes, error) {
	user := &domain.User{
		Name:     req.Name,
		Password: req.Password,
		Email:    req.Email,
	}
	err := u.userRepo.AddUser(ctx, user)
	return &v1.AddUserRes{Id: user.ID}, err
}

func (u UserRepoHandler) GetUserById(ctx context.Context, req *v1.UserIdReq) (*v1.UserRes, error) {
	user, err := u.userRepo.GetUserById(ctx, req.Id)
	res := &v1.UserRes{}
	if err != nil {
		return res, err
	}
	err = copier.Copy(res, user)
	return res, err
}

func (u UserRepoHandler) UserList(ctx context.Context, req *v1.UserListReq) (*v1.UserListRes, error) {
	res := &v1.UserListRes{}
	list, err := u.userRepo.GetUserList(ctx, req.Ids)
	if err != nil {
		return res, nil
	}
	userList := make([]*v1.UserRes, 0, len(list))
	err = copier.Copy(userList, list)
	res.UserList = userList
	return res, err
}

func (u UserRepoHandler) GetByEmail(ctx context.Context,req *v1.EmailReq) (*v1.UserRes, error) {
	user, err := u.userRepo.GetUserByEmail(ctx, req.Email)
	res := &v1.UserRes{}
	if err != nil {
		return res, err
	}
	err = copier.Copy(res, user)
	return res, err
}

func (u UserRepoHandler) mustEmbedUnimplementedUserRepoServiceServer() {
}
