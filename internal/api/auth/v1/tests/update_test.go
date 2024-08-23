package tests

import (
	"context"
	"testing"

	server "github.com/andredubov/auth/internal/api/auth/v1"
	"github.com/andredubov/auth/internal/service"
	serviceMocks "github.com/andredubov/auth/internal/service/mocks"
	"github.com/andredubov/auth/internal/service/model"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dvln/testify/require"
	"github.com/gojuno/minimock/v3"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdateUser(t *testing.T) {
	t.Parallel()
	type usersServiceMockFunc func(mc *minimock.Controller) service.Users
	type args struct {
		ctx context.Context
		req *auth_v1.UpdateRequest
	}

	var (
		ctx            = context.Background()
		mc             = minimock.NewController(t)
		id             = gofakeit.Int64()
		name           = gofakeit.Name()
		email          = gofakeit.Email()
		role           = model.Role(gofakeit.IntRange(1, 2))
		serviceError   = status.Error(codes.Internal, "failed to update user")
		updateUserInfo = model.UpdateUserInfo{
			ID:    id,
			Name:  &name,
			Email: &email,
			Role:  &role,
		}
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
				req: &auth_v1.UpdateRequest{
					Id: id,
					Info: &auth_v1.UpdateUserInfo{
						Name:  wrapperspb.String(name),
						Email: wrapperspb.String(email),
						Role:  auth_v1.UserRole(role),
					},
				},
			},
			want: &empty.Empty{},
			err:  nil,
			usersServiceMock: func(mc *minimock.Controller) service.Users {
				mock := serviceMocks.NewUsersMock(mc)
				mock.UpdateMock.Expect(ctx, updateUserInfo).Return(id, nil)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: &auth_v1.UpdateRequest{
					Id: id,
					Info: &auth_v1.UpdateUserInfo{
						Name:  wrapperspb.String(name),
						Email: wrapperspb.String(email),
						Role:  auth_v1.UserRole(role),
					},
				},
			},
			want: nil,
			err:  serviceError,
			usersServiceMock: func(mc *minimock.Controller) service.Users {
				mock := serviceMocks.NewUsersMock(mc)
				mock.UpdateMock.Expect(ctx, updateUserInfo).Return(0, serviceError)
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

			newID, err := api.Update(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
