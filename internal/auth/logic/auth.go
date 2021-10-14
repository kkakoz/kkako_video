package logic

import (
	"context"
	"github.com/pkg/errors"
	v1 "kkako_video/api/user/v1"
	"kkako_video/internal/auth/domain"
	"kkako_video/internal/pkg/jwtx"
	"time"
)

type AuthLogic struct {
	userCli  v1.UserServiceClient
	authRepo domain.IAuthRepo
	jwtGen   *jwtx.JwtTokenGen
}

func (u AuthLogic) Register(ctx context.Context, auth *domain.Auth, name string) error {
	auth, err := u.authRepo.GetAuthByEmail(ctx, auth.Email)
	if err != nil {
		return err
	}
	if auth.ID != 0 {
		return errors.New("该邮件已经注册")
	}
	err = u.authRepo.AddAuth(ctx, auth)
	if err != nil {
		return err
	}
	_, err = u.userCli.AddUser(ctx, &v1.AddUserReq{Name: name})
	if err != nil {
		e := u.authRepo.DeleteAuth(ctx, auth.ID)
		if e != nil {
			return e
		}
		return err
	}
	return nil
}

func (u AuthLogic) Login(ctx context.Context, auth *domain.Auth) (int64, string, error) {
	auth, err := u.authRepo.GetAuthByEmail(ctx, auth.Email)
	if err != nil {
		return 0, "", err
	}
	user, err := u.userCli.GetUser(ctx, &v1.GetUserReq{Id: auth.UserId})
	if err != nil {
		return 0, "", err
	}
	token, err := u.jwtGen.GenTokenExpire(auth.UserId, user.Name, user.State, auth.Auth, time.Hour*24*3)
	if err != nil {
		return 0, "", err
	}
	return auth.UserId, token, nil
}

func NewAuthLogic(userCli v1.UserServiceClient, authRepo domain.IAuthRepo, jwtGen *jwtx.JwtTokenGen) domain.IAuthLogic {
	return &AuthLogic{userCli: userCli, authRepo: authRepo, jwtGen: jwtGen}
}
