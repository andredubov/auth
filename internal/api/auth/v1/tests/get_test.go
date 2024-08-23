package tests

import (
	"context"
	"testing"
	"time"

	server "github.com/andredubov/auth/internal/api/auth/v1"
	"github.com/andredubov/auth/internal/service"
	"github.com/andredubov/auth/internal/service/converter"
	serviceMocks "github.com/andredubov/auth/internal/service/mocks"
	"github.com/andredubov/auth/internal/service/model"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dvln/testify/require"
	"github.com/gojuno/minimock/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestGetUser(t *testing.T) {
	t.Parallel()
	type usersServiceMockFunc func(mc *minimock.Controller) service.Users
	type args struct {
		ctx context.Context
		req *auth_v1.GetRequest
	}

	var (
		ctx          = context.Background()
		mc           = minimock.NewController(t)
		id           = gofakeit.Int64()
		name         = gofakeit.Name()
		email        = gofakeit.Email()
		role         = gofakeit.IntRange(1, 2)
		serviceError = status.Error(codes.Internal, "failed to get user")

		user = &model.User{
			ID:        id,
			Name:      name,
			Email:     email,
			Role:      model.Role(role),
			CreatedAt: time.Now().UTC(),
		}
	)

	tests := []struct {
		name             string
		args             args
		want             *auth_v1.GetResponse
		err              error
		usersServiceMock usersServiceMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: &auth_v1.GetRequest{
					Id: id,
				},
			},
			want: &auth_v1.GetResponse{
				User: converter.ToUserFromService(user),
			},
			err: nil,
			usersServiceMock: func(mc *minimock.Controller) service.Users {
				mock := serviceMocks.NewUsersMock(mc)
				mock.GetByIDMock.Expect(ctx, id).Return(user, nil)
				return mock
			},
		},
		{
			name: "error case",
			args: args{
				ctx: ctx,
				req: &auth_v1.GetRequest{
					Id: id,
				},
			},
			want: nil,
			err:  serviceError,
			usersServiceMock: func(mc *minimock.Controller) service.Users {
				mock := serviceMocks.NewUsersMock(mc)
				mock.GetByIDMock.Expect(ctx, id).Return(nil, serviceError)
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

			newID, err := api.Get(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
