package handler

import (
	"context"
	"github.com/jinzhu/copier"
	v1 "kkako_video/api/user/v1"
	"kkako_video/internal/user_repo/domain"
	"kkako_video/pkg/cryption"
	"kkako_video/pkg/db/mysqlx"
)

type UserRepoHandler struct {
	v1.UnimplementedUserRepoServiceServer
	userRepo domain.IUserRepo
}

func NewUserRepoHandler(userRepo domain.IUserRepo) *UserRepoHandler {
	db := mysqlx.GetDB(context.TODO())
	db.AutoMigrate(&domain.User{})
	return &UserRepoHandler{userRepo: userRepo}
}

func (u UserRepoHandler) AddUser(ctx context.Context, req *v1.AddUserReq) (*v1.AddUserRes, error) {
	user := &domain.User{
		Name:     req.Name,
		Email:    req.Email,
	}
	user.Salt = cryption.UUID()
	user.Password = cryption.Md5Str(req.Password + user.Salt)
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
	for _, user := range list {
		userRes := &v1.UserRes{}
		err = copier.Copy(userRes, user)
		if err != nil {
			return nil, err
		}
		userList = append(userList, userRes)
	}
	res.UserList = userList
	return res, err
}

func (u UserRepoHandler) GetByCondition(ctx context.Context,req *v1.ConditionReq) (*v1.UserRes, error) {
	user, err := u.userRepo.GetUser(ctx, req.Email)
	res := &v1.UserRes{}
	if err != nil {
		return res, err
	}
	err = copier.Copy(res, user)
	return res, err
}

func (u UserRepoHandler) UpdateUser(context.Context, *v1.UserUpdateReq) (*v1.UserUpdateRes, error) {
	panic("un implement")
}