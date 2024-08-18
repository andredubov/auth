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

func TestGetUser(t *testing.T) {
	t.Parallel()
	type passwordHasherMockFunc func(mc *minimock.Controller) hasher.PasswordHasher
	type usersRepositoryMockFunc func(mc *minimock.Controller) repository.Users
	type usersCacheMockFunc func(mc *minimock.Controller) cache.Users
	type txManagerMockFunc func(mc *minimock.Controller) database.TxManager

	type args struct {
		ctx    context.Context
		userID int64
		output *model.User
	}

	var (
		ctx             = context.Background()
		mc              = minimock.NewController(t)
		repositoryError = errors.New("repo error")
		cacheError      = errors.New("cache error")
		foundUser       = &model.User{
			ID:       gofakeit.Int64(),
			Name:     gofakeit.Name(),
			Email:    gofakeit.Email(),
			Role:     model.Role(gofakeit.IntRange(1, 2)),
			Password: gofakeit.Password(true, true, true, true, false, 10),
		}
		userID = gofakeit.Int64()
	)

	tests := []struct {
		name                string
		args                args
		want                *model.User
		err                 error
		passwordHasherMock  passwordHasherMockFunc
		usersRepositoryMock usersRepositoryMockFunc
		usersCacheMock      usersCacheMockFunc
		txManagerMock       txManagerMockFunc
	}{
		{
			name: "success case",
			args: args{
				ctx:    ctx,
				userID: userID,
				output: foundUser,
			},
			want: foundUser,
			err:  nil,
			passwordHasherMock: func(mc *minimock.Controller) hasher.PasswordHasher {
				mock := hasherMocks.NewPasswordHasherMock(mc)
				return mock
			},
			usersRepositoryMock: func(mc *minimock.Controller) repository.Users {
				mock := repoMocks.NewUsersMock(mc)
				return mock
			},
			usersCacheMock: func(mc *minimock.Controller) cache.Users {
				mock := cacheMocks.NewUsersMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(foundUser, nil)
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
		{
			name: "success case from repo",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want: foundUser,
			err:  nil,
			passwordHasherMock: func(mc *minimock.Controller) hasher.PasswordHasher {
				mock := hasherMocks.NewPasswordHasherMock(mc)
				return mock
			},
			usersRepositoryMock: func(mc *minimock.Controller) repository.Users {
				mock := repoMocks.NewUsersMock(mc)
				mock.GetByIDMock.Expect(ctx, userID).Return(foundUser, nil)
				return mock
			},
			usersCacheMock: func(mc *minimock.Controller) cache.Users {
				mock := cacheMocks.NewUsersMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(nil, cacheError)
				mock.CreateMock.Expect(ctx, *foundUser).Return(nil)
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
		{
			name: "service error case (get user repo error)",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want: nil,
			err:  repositoryError,
			passwordHasherMock: func(mc *minimock.Controller) hasher.PasswordHasher {
				mock := hasherMocks.NewPasswordHasherMock(mc)
				return mock
			},
			usersRepositoryMock: func(mc *minimock.Controller) repository.Users {
				mock := repoMocks.NewUsersMock(mc)
				mock.GetByIDMock.Expect(ctx, userID).Return(nil, repositoryError)
				return mock
			},
			usersCacheMock: func(mc *minimock.Controller) cache.Users {
				mock := cacheMocks.NewUsersMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(nil, cacheError)
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
		{
			name: "service error case (create cache error)",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want: nil,
			err:  cacheError,
			passwordHasherMock: func(mc *minimock.Controller) hasher.PasswordHasher {
				mock := hasherMocks.NewPasswordHasherMock(mc)
				return mock
			},
			usersRepositoryMock: func(mc *minimock.Controller) repository.Users {
				mock := repoMocks.NewUsersMock(mc)
				mock.GetByIDMock.Expect(ctx, userID).Return(foundUser, nil)
				return mock
			},
			usersCacheMock: func(mc *minimock.Controller) cache.Users {
				mock := cacheMocks.NewUsersMock(mc)
				mock.GetMock.Expect(ctx, userID).Return(nil, cacheError)
				mock.CreateMock.Expect(ctx, *foundUser).Return(cacheError)
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

			targetUser, err := service.GetByID(tt.args.ctx, tt.args.userID)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, targetUser)
		})
	}
}
