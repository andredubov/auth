package tests

import (
	"context"
	"errors"
	"testing"

	"github.com/andredubov/auth/internal/cache"
	cacheMocks "github.com/andredubov/auth/internal/cache/mocks"
	"github.com/andredubov/auth/internal/repository"
	repoMocks "github.com/andredubov/auth/internal/repository/mocks"
	"github.com/andredubov/auth/internal/service/user"
	"github.com/andredubov/golibs/pkg/client/database"
	dbMocks "github.com/andredubov/golibs/pkg/client/database/mocks"
	"github.com/andredubov/golibs/pkg/hasher"
	hasherMocks "github.com/andredubov/golibs/pkg/hasher/mocks"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/dvln/testify/require"
	"github.com/gojuno/minimock/v3"
)

func TestDeleteUser(t *testing.T) {
	t.Parallel()
	type passwordHasherMockFunc func(mc *minimock.Controller) hasher.PasswordHasher
	type usersRepositoryMockFunc func(mc *minimock.Controller) repository.Users
	type usersCacheMockFunc func(mc *minimock.Controller) cache.Users
	type txManagerMockFunc func(mc *minimock.Controller) database.TxManager

	type args struct {
		ctx    context.Context
		userID int64
	}

	var (
		ctx             = context.Background()
		mc              = minimock.NewController(t)
		repositoryError = errors.New("repo error")
		userID          = gofakeit.Int64()
		rowsAffected    = int64(1)
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
				ctx:    ctx,
				userID: userID,
			},
			want: rowsAffected,
			err:  nil,
			passwordHasherMock: func(mc *minimock.Controller) hasher.PasswordHasher {
				mock := hasherMocks.NewPasswordHasherMock(mc)
				return mock
			},
			usersRepositoryMock: func(mc *minimock.Controller) repository.Users {
				mock := repoMocks.NewUsersMock(mc)
				mock.DeleteMock.Expect(ctx, userID).Return(rowsAffected, nil)
				return mock
			},
			usersCacheMock: func(mc *minimock.Controller) cache.Users {
				mock := cacheMocks.NewUsersMock(mc)
				mock.DeleteMock.Expect(ctx, userID).Return(nil)
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
			name: "service error case",
			args: args{
				ctx:    ctx,
				userID: userID,
			},
			want: 0,
			err:  repositoryError,
			passwordHasherMock: func(mc *minimock.Controller) hasher.PasswordHasher {
				mock := hasherMocks.NewPasswordHasherMock(mc)
				return mock
			},
			usersRepositoryMock: func(mc *minimock.Controller) repository.Users {
				mock := repoMocks.NewUsersMock(mc)
				mock.DeleteMock.Expect(ctx, userID).Return(0, repositoryError)
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

			result, err := service.Delete(tt.args.ctx, tt.args.userID)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
