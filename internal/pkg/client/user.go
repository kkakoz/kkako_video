package client

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	v1 "kkako_video/api/user/v1"
)

func NewUserClient() (v1.UserServiceClient, error) {
	conn, err := grpc.Dial("dns///" + UserAddr, grpc.WithInsecure(), grpc.WithBalancerName(roundrobin.Name))
	if err != nil {
		return nil, err
	}
	client := v1.NewUserServiceClient(conn)
	return client, nil
}

