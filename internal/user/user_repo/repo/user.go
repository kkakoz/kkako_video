package repo

import (
	"context"
	"github.com/pkg/errors"
	"kkako_video/internal/user/user_repo/domain"
	"kkako_video/pkg/db/mysqlx"
)

type UserRepo struct {

}

func NewUserRepo() domain.IUserRepo {
	return &UserRepo{}
}

func (u UserRepo) GetUserByEmail(ctx context.Context, email string) (user *domain.User, err error) {
	db := mysqlx.GetDB(ctx)
	err = db.Where("email = ?", email).Find(user).Error
	return user, errors.Wrap(err, "添加用户失败")
}

func (u UserRepo) AddUser(ctx context.Context, user *domain.User) error {
	db := mysqlx.GetDB(ctx)
	err := db.Create(user).Error
	return errors.Wrap(err, "添加用户失败")
}

func (u UserRepo) GetUserById(ctx context.Context, id int64) (user *domain.User, err error) {
	db := mysqlx.GetDB(ctx)
	err = db.Where("id = ?", id).Find(user).Error
	return user, errors.Wrap(err, "查找失败")
}

func (u UserRepo) GetUserList(ctx context.Context, ids []int64) (list []*domain.User, err error) {
	db := mysqlx.GetDB(ctx)
	err = db.Where("id in ?", ids).Find(&list).Error
	return list, errors.Wrap(err, "查找失败")
}

