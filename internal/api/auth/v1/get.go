package server

import (
	"context"

	"github.com/andredubov/auth/internal/service/converter"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Get a user
func (i *Implementation) Get(ctx context.Context, r *auth_v1.GetRequest) (*auth_v1.GetResponse, error) {
	user, err := i.usersService.GetByID(ctx, r.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	return &auth_v1.GetResponse{
		User: converter.ToUserFromService(user),
	}, nil
}
