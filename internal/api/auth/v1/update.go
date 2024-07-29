package server

import (
	"context"

	"github.com/andredubov/auth/internal/service/converter"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *Implementation) Update(ctx context.Context, r *auth_v1.UpdateRequest) (*empty.Empty, error) {
	if _, err := i.usersService.Update(ctx, converter.ToUserUpdateInfoFromUpdateRequest(r)); err != nil {
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	return &empty.Empty{}, nil
}
