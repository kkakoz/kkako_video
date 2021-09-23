package handler_test

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/spf13/viper"
	v1 "kkako_video/api/user/v1"
	"kkako_video/internal/user_repo/handler"
	"kkako_video/internal/user_repo/repo"
	"kkako_video/pkg/cryption"
	"kkako_video/pkg/db/mysqlx"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	viper.AddConfigPath("../../../config")
	viper.SetConfigName("test")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln("read conf err:", err)
	}
	_, err = mysqlx.New(viper.GetViper())
	if err != nil {
		log.Fatalln(err)
	}
	mysqlx.FlushDB()
	m.Run()
}

func TestUserRepoHandler(t *testing.T) {
	Convey("user_repo", t, func() {
		repoHandler := handler.NewUserRepoHandler(repo.NewUserRepo())
		ctx := context.TODO()
		testData := []*v1.AddUserReq{
				{
					Name:     "zhangsan",
					Email:    "zhangsan@qq.com",
					Password: "123456",
				},
				{
					Name:     "lisi",
					Email:    "lisi@qq.com",
					Password: "password",
				},
				{
					Name:     "lisi",
					Email:    "lisi2@qq.com",
					Password: "password2",
				},
		}
		ids := make([]int64, 0, len(testData))

		for _, req := range testData {
			id, err := repoHandler.AddUser(ctx, req)
			So(err, ShouldBeNil)
			ids = append(ids, id.Id)
		}

		for i, id := range ids {
			req := &v1.UserIdReq{Id: id}
			user, err := repoHandler.GetUserById(ctx, req)
			user1 := testData[i]
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)
			So(user.Name, ShouldEqual, user1.Name)
			So(user.Email, ShouldEqual, user1.Email)
			So(user.Password, ShouldEqual, cryption.Md5Str(user1.Password + user.Salt))
		}

		req := &v1.UserListReq{Ids: ids}
		list, err := repoHandler.UserList(context.TODO(), req)
		So(err, ShouldBeNil)
		for i, user := range list.UserList {
			So(user.Name, ShouldEqual, testData[i].Name)
			So(user.Email, ShouldEqual, testData[i].Email)
		}

	})
}
