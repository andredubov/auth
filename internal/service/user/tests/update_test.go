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

func TestUpdateUser(t *testing.T) {
	t.Parallel()
	type passwordHasherMockFunc func(mc *minimock.Controller) hasher.PasswordHasher
	type usersRepositoryMockFunc func(mc *minimock.Controller) repository.Users
	type usersCacheMockFunc func(mc *minimock.Controller) cache.Users
	type txManagerMockFunc func(mc *minimock.Controller) database.TxManager

	type args struct {
		ctx   context.Context
		input model.UpdateUserInfo
	}

	var (
		ctx             = context.Background()
		mc              = minimock.NewController(t)
		repositoryError = errors.New("repo error")
		rowsAffected    = int64(1)
		name            = gofakeit.Name()
		email           = gofakeit.Email()
		role            = model.Role(gofakeit.IntRange(1, 2))
		updateUserInfo  = model.UpdateUserInfo{
			ID:    gofakeit.Int64(),
			Name:  &name,
			Email: &email,
			Role:  &role,
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
				ctx:   ctx,
				input: updateUserInfo,
			},
			want: rowsAffected,
			err:  nil,
			passwordHasherMock: func(mc *minimock.Controller) hasher.PasswordHasher {
				mock := hasherMocks.NewPasswordHasherMock(mc)
				return mock
			},
			usersRepositoryMock: func(mc *minimock.Controller) repository.Users {
				mock := repoMocks.NewUsersMock(mc)
				mock.UpdateMock.Expect(ctx, updateUserInfo).Return(rowsAffected, nil)
				return mock
			},
			usersCacheMock: func(mc *minimock.Controller) cache.Users {
				mock := cacheMocks.NewUsersMock(mc)
				mock.DeleteMock.Expect(ctx, updateUserInfo.ID).Return(nil)
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
				ctx:   ctx,
				input: updateUserInfo,
			},
			want: 0,
			err:  repositoryError,
			passwordHasherMock: func(mc *minimock.Controller) hasher.PasswordHasher {
				mock := hasherMocks.NewPasswordHasherMock(mc)
				return mock
			},
			usersRepositoryMock: func(mc *minimock.Controller) repository.Users {
				mock := repoMocks.NewUsersMock(mc)
				mock.UpdateMock.Expect(ctx, updateUserInfo).Return(0, repositoryError)
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

			result, err := service.Update(tt.args.ctx, tt.args.input)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, result)
		})
	}
}
