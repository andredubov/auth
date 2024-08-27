package tests

import (
	"context"
	"testing"

	server "github.com/andredubov/auth/internal/api/auth/v1"
	"github.com/andredubov/auth/internal/service"
	serviceMocks "github.com/andredubov/auth/internal/service/mocks"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dvln/testify/require"
	"github.com/gojuno/minimock/v3"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestDeleteUser(t *testing.T) {
	t.Parallel()
	type usersServiceMockFunc func(mc *minimock.Controller) service.Users
	type args struct {
		ctx context.Context
		req *auth_v1.DeleteRequest
	}

	var (
		ctx          = context.Background()
		mc           = minimock.NewController(t)
		id           = gofakeit.Int64()
		serviceError = status.Error(codes.Internal, "failed to delete user")
	)

	tests := []struct {
		name             string
		args             args
		want             *empty.Empty
		err              error
		usersServiceMock usersServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: &auth_v1.DeleteRequest{
					Id: id,
				},
			},
			want: &empty.Empty{},
			err:  nil,
			usersServiceMock: func(mc *minimock.Controller) service.Users {
				mock := serviceMocks.NewUsersMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(id, nil)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: &auth_v1.DeleteRequest{
					Id: id,
				},
			},
			want: nil,
			err:  serviceError,
			usersServiceMock: func(mc *minimock.Controller) service.Users {
				mock := serviceMocks.NewUsersMock(mc)
				mock.DeleteMock.Expect(ctx, id).Return(0, serviceError)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			usersServiceMock := tt.usersServiceMock(mc)
			api := server.NewImplementation(usersServiceMock)

			newID, err := api.Delete(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
