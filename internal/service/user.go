package service

import (
	"context"
	v1 "cube-core/api/realworld/v1"
)

func (r *RealWorldService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.UserReply, error) {
	return &v1.UserReply{User: &v1.UserReply_User{
		Username: "boom",
	}}, nil
}
