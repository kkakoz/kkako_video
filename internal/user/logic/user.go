package logic

import (
	"context"
	"kkako_video/internal/user/domain"
)

type UserLogic struct {
	userRepo domain.IUserRepo
}

func (u UserLogic) AddUser(ctx context.Context, user *domain.User) error {
	return u.userRepo.AddUser(ctx, user)
}

func (u UserLogic) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	return u.userRepo.GetUser(ctx, id)
}

func (u UserLogic) GetUsers(ctx context.Context, ids []int64) ([]*domain.User, error) {
	return u.userRepo.GetUserList(ctx, ids)
}

func NewUserLogic() domain.IUserLogic {
	return &UserLogic{}
}