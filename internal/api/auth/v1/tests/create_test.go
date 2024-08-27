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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	type usersServiceMockFunc func(mc *minimock.Controller) service.Users
	type args struct {
		ctx context.Context
		req *auth_v1.CreateRequest
	}

	var (
		ctx = context.Background()
		mc  = minimock.NewController(t)

		id            = gofakeit.Int64()
		name          = gofakeit.Name()
		email         = gofakeit.Email()
		role          = gofakeit.IntRange(1, 2)
		plainPassword = gofakeit.Password(true, true, true, true, false, 10)

		serviceError = status.Error(codes.Internal, "failed to create a new user")

		user = model.User{
			Name:            name,
			Email:           email,
			Role:            model.Role(role),
			Password:        plainPassword,
			PasswordConfirm: plainPassword,
		}

		response = &auth_v1.CreateResponse{
			Id: id,
		}
	)

	tests := []struct {
		name             string
		args             args
		want             *auth_v1.CreateResponse
		err              error
		usersServiceMock usersServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: &auth_v1.CreateRequest{
					Info: &auth_v1.UserInfo{
						Name:  name,
						Email: email,
						Role:  auth_v1.UserRole(role),
					},
					Password:        plainPassword,
					PasswordConfirm: plainPassword,
				},
			},
			want: response,
			err:  nil,
			usersServiceMock: func(mc *minimock.Controller) service.Users {
				mock := serviceMocks.NewUsersMock(mc)
				mock.CreateMock.Expect(ctx, user).Return(id, nil)
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: &auth_v1.CreateRequest{
					Info: &auth_v1.UserInfo{
						Name:  name,
						Email: email,
						Role:  auth_v1.UserRole(role),
					},
					Password:        plainPassword,
					PasswordConfirm: plainPassword,
				},
			},
			want: nil,
			err:  serviceError,
			usersServiceMock: func(mc *minimock.Controller) service.Users {
				mock := serviceMocks.NewUsersMock(mc)
				mock.CreateMock.Expect(ctx, user).Return(0, serviceError)
				return mock
			},
		},
		{
			name: "password empty",
			args: args{
				ctx: ctx,
				req: &auth_v1.CreateRequest{
					Info: &auth_v1.UserInfo{
						Name:  name,
						Email: email,
						Role:  auth_v1.UserRole(role),
					},
				},
			},
			want: nil,
			err:  status.Error(codes.InvalidArgument, "user password empty"),
			usersServiceMock: func(mc *minimock.Controller) service.Users {
				mock := serviceMocks.NewUsersMock(mc)
				return mock
			},
		},
		{
			name: "password and password confirm do not match",
			args: args{
				ctx: ctx,
				req: &auth_v1.CreateRequest{
					Info: &auth_v1.UserInfo{
						Name:  name,
						Email: email,
						Role:  auth_v1.UserRole(role),
					},
					Password:        plainPassword,
					PasswordConfirm: plainPassword + "a",
				},
			},
			want: nil,
			err:  status.Error(codes.InvalidArgument, "password doesn't equal to password_confirm"),
			usersServiceMock: func(mc *minimock.Controller) service.Users {
				mock := serviceMocks.NewUsersMock(mc)
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

			newID, err := api.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
