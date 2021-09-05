package client

import (
	"google.golang.org/grpc"
	v1 "kkako_video/api/user/v1"
)

func NewUserRepoClient() v1.UserRepoServiceClient {
	conn, err := grpc.Dial("dns///")
	if err != nil {

	}
	client := v1.NewUserRepoServiceClient(conn)
	return client
}
