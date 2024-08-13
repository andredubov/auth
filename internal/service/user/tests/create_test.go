package tests

import (
	"context"
	"fmt"
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
	"github.com/brianvoe/gofakeit"
	"github.com/dvln/testify/require"
	"github.com/gojuno/minimock/v3"
)

func TestCreateUser(t *testing.T) {
	t.Parallel()
	type usersRepositoryMockFunc func(mc *minimock.Controller) repository.Users
	type passwordHasherMockFunc func(mc *minimock.Controller) hasher.PasswordHasher
	type usersCacheMockFunc func(mc *minimock.Controller) cache.Users
	type txManagerMockFunc func(f func(context.Context) error, mc *minimock.Controller) database.TxManager

	type args struct {
		ctx context.Context
		req model.User
	}

	var (
		ctx             = context.Background()
		controller      = minimock.NewController(t)
		repositoryError = fmt.Errorf("repo error")
		id              = gofakeit.Int64()
		name            = gofakeit.Name()
		email           = gofakeit.Email()
		role            = model.Role(1)
		plainPassword   = gofakeit.Password(true, true, true, true, false, 10)
		hashedPassword  = gofakeit.Password(true, true, true, true, false, 10)

		request = model.User{
			ID:              id,
			Name:            name,
			Email:           email,
			Role:            role,
			Password:        plainPassword,
			PasswordConfirm: plainPassword,
		}

		cacheUser = model.User{
			ID:    id,
			Name:  name,
			Email: email,
			Role:  role,
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
				mock.CreateMock.Expect(ctx, request).Return(id, nil)
				return mock
			},
			usersCacheMock: func(mc *minimock.Controller) cache.Users {
				mock := cacheMocks.NewUsersMock(mc)
				mock.CreateMock.Expect(ctx, &cacheUser).Return(nil)
				return mock
			},
			txManagerMock: func(f func(context.Context) error, mc *minimock.Controller) database.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Expect(ctx, f).Return(nil)
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
				mock.CreateMock.Expect(ctx, request).Return(0, repositoryError)
				return mock
			},
			usersCacheMock: func(mc *minimock.Controller) cache.Users {
				mock := cacheMocks.NewUsersMock(mc)
				mock.CreateMock.Expect(ctx, &cacheUser).Return(nil)
				return mock
			},
			txManagerMock: func(f func(context.Context) error, mc *minimock.Controller) database.TxManager {
				mock := dbMocks.NewTxManagerMock(mc)
				mock.ReadCommittedMock.Expect(ctx, f).Return(nil)
				return mock
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			usersRepositoryMock := tt.usersRepositoryMock(controller)
			usersCacheMock := tt.usersCacheMock(controller)
			passwordHasherMock := tt.passwordHasherMock(controller)
			txManagerMock := tt.txManagerMock(func(ctx context.Context) error {
				var errTx error
				hashedPassword, errTx := passwordHasherMock.HashAndSalt(request.Password)
				if errTx != nil {
					return errTx
				}

				request.Password = hashedPassword

				id, errTx = usersRepositoryMock.Create(ctx, request)
				if errTx != nil {
					return errTx
				}

				errTx = usersCacheMock.Create(ctx, &cacheUser)
				if errTx != nil {
					return errTx
				}

				return nil
			}, controller)

			service := user.NewService(
				usersRepositoryMock,
				passwordHasherMock,
				usersCacheMock,
				txManagerMock,
			)

			id, err := service.Create(tt.args.ctx, tt.args.req)
			require.Equal(t, tt.err, err)
			require.Equal(t, tt.want, id)
		})
	}
}
