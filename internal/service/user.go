package service

import (
	"context"
	v1 "cube-core/api/realworld/v1"
)

func (r *RealWorldService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.UserReply, error) {
	rv, err := r.uc.Login(ctx, request.User.Email, request.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{User: &v1.UserReply_User{
		Username: rv.Username,
		Token:    rv.Token,
	}}, nil
}

func (r *RealWorldService) Register(ctx context.Context, request *v1.RegisterRequest) (*v1.UserReply, error) {
	u, err := r.uc.Register(ctx, request.User.Username, request.User.Email, request.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{User: &v1.UserReply_User{
		Email:    u.Email,
		Username: u.Username,
		Token:    u.Token,
	}}, nil
}

func (r *RealWorldService) GetCurrentUser(ctx context.Context, request *v1.GetCurrentUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{User: &v1.UserReply_User{
		Username: "boom",
	}}, nil
}
func (r *RealWorldService) UpdateUser(ctx context.Context, request *v1.UpdateUserRequest) (*v1.UserReply, error) {
	return &v1.UserReply{User: &v1.UserReply_User{
		Username: "boom",
	}}, nil
}
