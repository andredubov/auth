package server

import (
	"context"

	"github.com/andredubov/auth/internal/service/converter"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Create(ctx context.Context, r *auth_v1.CreateRequest) (*auth_v1.CreateResponse, error) {
	if r.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "user password empty")
	}

	if r.GetPassword() != r.GetPasswordConfirm() {
		return nil, status.Error(codes.InvalidArgument, "password doesn't equal to password_confirm")
	}

	id, err := i.usersService.Create(ctx, converter.ToUserFromCreateRequest(r))
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to create a new user")
	}

	return &auth_v1.CreateResponse{
		Id: id,
	}, nil
}
