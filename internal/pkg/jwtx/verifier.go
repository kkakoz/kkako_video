package jwtx

import (
	"context"
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
)

type JwtTokenVerifier struct {
	PublicKeyPath string
	publicKey     *rsa.PublicKey
}

func NewJwtTokenVerifier(viper *viper.Viper) (*JwtTokenVerifier, error) {
	verifier := &JwtTokenVerifier{}
	err := viper.UnmarshalKey("jwt", verifier)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(verifier.PublicKeyPath)
	if err != nil {
		return nil, errors.Wrap(err, "打开rsa privatekey文件失败")
	}
	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, errors.Wrap(err, "读取文件失败")
	}
	pem, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		return nil, err
	}
	return &JwtTokenVerifier{publicKey: pem}, nil
}

func (v *JwtTokenVerifier) Verifier(token string) (*UserClaims, error) {
	t, err := jwt.ParseWithClaims(token, &UserClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return v.publicKey, nil
		})
	if err != nil {
		return nil, err
	}
	if !t.Valid {
		return nil, err
	}
	clm, ok := t.Claims.(*UserClaims)
	if !ok {
		return nil, err
	}

	if err = clm.Valid(); err != nil {
		return nil, err
	}

	return clm, nil
}

var userClaimsKey struct{}

func WithUser(ctx context.Context, claims *UserClaims) context.Context {
	return context.WithValue(ctx, userClaimsKey, claims)
}

func GetUser(ctx context.Context) (*UserClaims, error) {
	value := ctx.Value(userClaimsKey)
	claims, ok := value.(*UserClaims)
	if !ok {
		return claims, errors.New("获取用户失败")
	}
	return claims, nil
}