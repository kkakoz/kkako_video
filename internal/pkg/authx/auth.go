package authx

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/pkg/errors"
	"kkako_video/internal/pkg/handle"
	"net/http"
	"strings"
	"time"
)

type User struct {
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Avatar      string  `json:"avatar"`
	Brief       string  `json:"brief"`
	FollowCount int64   `json:"follow_count"`
	FansCount   int64   `json:"fans_count"`
	LikeCount   int64   `json:"like_count"`
	State       int32   `json:"state"`
}

func GetUserFromCtx(ctx context.Context) (*User, error) {
	value, ok := ctx.Value(CtxCurrentUserKey).(string)
	if !ok {
		return nil, errors.New("获取当前用户失败")
	}
	user := &User{}
	err := json.Unmarshal([]byte(value), user)
	if err != nil {
		return user, err
	}
	return nil, nil
}

func SetUser(client *redis.Client, user *User, token string) error {
	err := client.Set(GetTokenPreKey(token), user, time.Hour*24*3).Err()
	if err != nil {
		return err
	}
	return client.SAdd(GetAuthTokens(user.ID), token).Err()
}

func AuthValidate(handler http.Handler, client redis.Client) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if !strings.Contains(request.RequestURI, "login") && !strings.Contains(request.RequestURI, "register") {
			token := request.Header.Get("authorization")
			result, err := client.Get(token).Result()
			if err != nil {
				res := handle.ErrHandler(err)
				data, _ := json.Marshal(res)
				writer.Write(data)
				return
			}
			request.Header.Set(CtxCurrentUserKey, result)
		}
		handler.ServeHTTP(writer, request)
	})
}
