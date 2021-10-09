package repo

import (
	"context"
	"github.com/jinzhu/copier"
	v1 "kkako_video/api/user/v1"
	"kkako_video/internal/user/domain"
)

type UserRepo struct {
	userRepo v1.UserRepoServiceClient
}

func (u UserRepo) Get(ctx context.Context, id int64) (*domain.User, error) {
	user, err := u.userRepo.GetUserById(ctx, &v1.UserIdReq{Id: id})
	if err != nil {
		return nil, err
	}
	res := &domain.User{}
	err = copier.Copy(res, user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (u UserRepo) GetList(ctx context.Context, ids []int64) ([]*domain.User, error) {
	panic("implement me")
}

func NewUserRepo(userRepo v1.UserRepoServiceClient) *UserRepo {
	return &UserRepo{userRepo: userRepo}
}

