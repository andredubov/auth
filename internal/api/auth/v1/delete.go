package server

import (
	"context"

	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Delete(ctx context.Context, r *auth_v1.DeleteRequest) (*empty.Empty, error) {
	if _, err := i.usersService.Delete(ctx, r.GetId()); err != nil {
		return nil, status.Error(codes.Internal, "failed to delete user")
	}

	return nil, nil
}
