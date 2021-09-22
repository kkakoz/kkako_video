package handler_test

import (
	"context"
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	v1 "kkako_video/api/user/v1"
	handler2 "kkako_video/internal/user_repo/handler"
	repo2 "kkako_video/internal/user_repo/repo"
	"kkako_video/pkg/db/mysqlx"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	_, err := mysqlx.New(viper.GetViper())
	if err != nil {
		log.Fatalln(err)
	}
	m.Run()
}

func TestUserRepoHandler(t *testing.T) {
	Convey("user_repo", t, func() {
		repoHandler := handler2.NewUserRepoHandler(repo2.NewUserRepo())
		ctx := context.TODO()
		type user struct {
			ID       int64
			Name     string
			Email    string
			Password string
		}
		//testData := []struct {
		//	req *v1.AddUserReq
		//	res user
		//}{
		//	{
		//		req: &v1.AddUserReq{},
		//		res: user{
		//			ID:       0,
		//			Name:     "",
		//			Email:    "",
		//			Password: "",
		//		},
		//	},
		//}
		Convey("add user", func() {
			req := &v1.AddUserReq{
				Name:     "testname",
				Email:    "test@163.com",
				Password: "test",
			}
			res, err := repoHandler.AddUser(ctx, req)
			So(err, ShouldBeNil)
			fmt.Println(res.Id)
		})
		Convey("user by id", func() {
			//req := &v1.UserIdReq{Id: 1}
			//user, err := repoHandler.GetUserById(ctx, req)
			//So(err, ShouldBeNil)
		})
		Convey("user list", func() {

		})
	})
}
