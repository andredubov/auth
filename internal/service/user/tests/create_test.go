package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/andredubov/auth/internal/cache"
	cacheMocks "github.com/andredubov/auth/internal/cache/mocks"
	"github.com/andredubov/auth/internal/repository"
	repoMocks "github.com/andredubov/auth/internal/repository/mocks"
	"github.com/andredubov/auth/internal/service/model"
	"github.com/andredubov/auth/internal/service/user"
	"github.com/andredubov/golibs/pkg/client/database"
	dbMocks "github.com/andredubov/golibs/pkg/client/database/mocks"
	"github.com/andredubov/golibs/pkg/hasher"
	hasherMocks "github.com/andredubov/golibs/pkg/hasher/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dvln/testify/require"
	"github.com/gojuno/minimock/v3"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	type usersRepositoryMockFunc func(mc *minimock.Controller) repository.Users
	type passwordHasherMockFunc func(mc *minimock.Controller) hasher.PasswordHasher
	type usersCacheMockFunc func(mc *minimock.Controller) cache.Users
	type txManagerMockFunc func(mc *minimock.Controller) database.TxManager

	type args struct {
		ctx context.Context
		req model.User
	}

	var (
		ctx             = context.Background()
		mc              = minimock.NewController(t)
		repositoryError = errors.New("repo error")
		id              = gofakeit.Int64()
		name            = gofakeit.Name()
		email           = gofakeit.Email()
		role            = model.Role(gofakeit.IntRange(1, 2))
		plainPassword   = gofakeit.Password(true, true, true, true, false, 10)
		hashedPassword  = gofakeit.Password(true, true, true, true, false, 10)

		request = model.User{
			Name:     name,
			Email:    email,
			Role:     role,
			Password: plainPassword,
		}

		repoUser = model.User{
			Name:     name,
			Email:    email,
			Role:     role,
			Password: hashedPassword,
		}

		cacheUser = model.User{
			ID:       id,
			Name:     name,
			Email:    email,
			Role:     role,
			Password: hashedPassword,
		}
	)

	tests := []struct {
		name                string
		args                args
		want                int64
		err                 error
		passwordHasherMock  passwordHasherMockFunc
		usersRepositoryMock usersRepositoryMockFunc
		usersCacheMock      usersCacheMockFunc
		txManagerMock       txManagerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx: ctx,
				req: request,
			},
			want: id,
			err:  nil,
			passwordHasherMock: func(mc *minimock.Controller) hasher.PasswordHasher {
				mock := hasherMocks.NewPasswordHasherMock(mc)
				mock.HashAndSaltMock.Expect(plainPassword).Return(hashedPassword, nil)
				return mock
			},
			usersRepositoryMock: func(mc *minimock.Controller) repository.Users {
				mock := repoMocks.NewUsersMock(mc)
				mock.CreateMock.Expect(ctx, repoUser).Return(id, nil)
				return mock
			},
			usersCacheMock: func(mc *minimock.Controller) cache.Users {
				mock := cacheMocks.NewUsersMock(mc)
				mock.CreateMock.Expect(ctx, cacheUser).Return(nil)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) database.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f database.Handler) (err error) {
					return f(ctx)
				})
				return mock
			},
		},
		{
			name: "service error case",
			args: args{
				ctx: ctx,
				req: request,
			},
			want: 0,
			err:  repositoryError,
			passwordHasherMock: func(mc *minimock.Controller) hasher.PasswordHasher {
				mock := hasherMocks.NewPasswordHasherMock(mc)
				mock.HashAndSaltMock.Expect(plainPassword).Return(hashedPassword, nil)
				return mock
			},
			usersRepositoryMock: func(mc *minimock.Controller) repository.Users {
				mock := repoMocks.NewUsersMock(mc)
				mock.CreateMock.Expect(ctx, repoUser).Return(0, repositoryError)
				return mock
			},
			usersCacheMock: func(mc *minimock.Controller) cache.Users {
				mock := cacheMocks.NewUsersMock(mc)
				return mock
			},
			txManagerMock: func(mc *minimock.Controller) database.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Set(func(ctx context.Context, f database.Handler) error {
					return f(ctx)
				})
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			passwordHasherMock := tt.passwordHasherMock(mc)
			usersRepositoryMock := tt.usersRepositoryMock(mc)
			usersCacheMock := tt.usersCacheMock(mc)
			txManagerMock := tt.txManagerMock(mc)

			service := user.NewService(
				usersRepositoryMock,
				passwordHasherMock,
				usersCacheMock,
				txManagerMock,
			)

			newID, err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, newID)
		})
	}
}
