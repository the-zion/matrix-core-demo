package service

import (
	"context"
	v1 "cube-core/api/realworld/v1"
	"cube-core/internal/errors"
)

func (r *RealWorldService) Login(ctx context.Context, request *v1.LoginRequest) (*v1.UserReply, error) {
	if len(request.User.Email) == 0 {
		return nil, errors.NewHTTPError(422, "email", "can not be empty")
	}
	return &v1.UserReply{User: &v1.UserReply_User{
		Username: "boom",
	}}, nil
}

func (r *RealWorldService) Register(ctx context.Context, request *v1.RegisterRequest) (*v1.UserReply, error) {
	u, err := r.uc.Register(ctx, request.User.Username, request.User.Email, request.User.Password)
	if err != nil {
		return nil, err
	}
	return &v1.UserReply{User: &v1.UserReply_User{
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
