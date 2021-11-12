package logic

import (
	"context"
	"github.com/go-redis/redis"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"kkako_video/internal/pkg/authx"
	"kkako_video/internal/user/domain"
	"kkako_video/pkg/cryption"
	"kkako_video/pkg/db/mysqlx"
)

var _ domain.IUserLogic = (*UserLogic)(nil)

func NewUserLogic(userRepo domain.IUserRepo, authRepo domain.IAuthRepo, redis *redis.Client) domain.IUserLogic {
	return &UserLogic{userRepo: userRepo, authRepo: authRepo, redis: redis}
}

type UserLogic struct {
	userRepo domain.IUserRepo
	authRepo domain.IAuthRepo
	redis    *redis.Client
}

func (u UserLogic) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	return u.userRepo.GetUser(ctx, id)
}

func (u UserLogic) GetUsers(ctx context.Context, ids []int64) ([]*domain.User, error) {
	return u.userRepo.GetUserList(ctx, ids)
}

func (u UserLogic) Register(ctx context.Context, auth *domain.Auth) (err error) {
	ctx, checkErr := mysqlx.Begin(ctx)
	defer func() {
		err = checkErr(err)
	}()
	preAuth, err := u.authRepo.GetAuthByEmail(ctx, auth.Email)
	if err != nil {
		return err
	}
	if preAuth.ID != 0 {
		return errors.New("该邮箱已经注册")
	}
	err = u.authRepo.AddAuth(ctx, auth)
	if err != nil {
		return err
	}
	user := &domain.User{}
	err = u.userRepo.AddUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u UserLogic) Login(ctx context.Context, auth *domain.Auth) (string, error) {
	a, err := u.authRepo.GetAuthByEmail(ctx, auth.Email)
	if err != nil {
		return "", err
	}
	user, err := u.userRepo.GetUser(ctx, a.UserId)
	if err != nil {
		return "", err
	}
	storeUser := &authx.User{}
	err = copier.Copy(storeUser, user)
	if err != nil {
		return "", err
	}
	token := cryption.UUID()
	err = authx.SetUser(u.redis, storeUser, token)
	return token, nil
}
