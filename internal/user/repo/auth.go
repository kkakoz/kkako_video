package repo

import (
	"context"
	"github.com/pkg/errors"
	"kkako_video/internal/user/domain"
	"kkako_video/pkg/db/mysqlx"
)

var _ domain.IAuthRepo = (*Auth)(nil)

func NewAuthRepo() domain.IAuthRepo {
	return &Auth{}
}

type Auth struct {

}

func (a Auth) GetAuthByEmail(ctx context.Context, email string) (*domain.Auth, error) {
	db := mysqlx.GetDB(ctx)
	user := &domain.Auth{}
	err := db.Where("email = ?", email).Find(user).Error
	return user, errors.Wrap(err, "查找失败")
}

func (a Auth) DeleteAuth(ctx context.Context, id int64) error {
	db := mysqlx.GetDB(ctx)
	err := db.Delete(&Auth{}, id).Error
	return errors.Wrap(err, "添加用户失败")
}

func (a Auth) AddAuth(ctx context.Context, user *domain.Auth) error {
	db := mysqlx.GetDB(ctx)
	err := db.Create(user).Error
	return errors.Wrap(err, "添加用户失败")
}

func (a Auth) GetAuth(ctx context.Context, id int64) (*domain.Auth, error) {
	db := mysqlx.GetDB(ctx)
	user := &domain.Auth{}
	err := db.Where("id = ?", id).Find(user).Error
	return user, errors.Wrap(err, "查找失败")
}
