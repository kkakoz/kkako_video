package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	v1 "kkako_video/api/user/v1"
)

func NewUserRepoClient() (v1.UserRepoServiceClient, error) {
	conn, err := grpc.Dial("dns///" + UserRepoAddr, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		return nil, err
	}
	client := v1.NewUserRepoServiceClient(conn)
	return client, nil
}
