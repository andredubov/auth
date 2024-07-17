package server

import (
	"context"
	"log"

	"github.com/andredubov/auth/internal/repository"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"github.com/andredubov/auth/pkg/hasher"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type authServer struct {
	auth_v1.UnimplementedAuthServer
	usersRepository repository.Users
	passwordHasher  hasher.PasswordHasher
}

func NewAuthServer(repository repository.Users, hasher hasher.PasswordHasher) auth_v1.AuthServer {
	return &authServer{
		usersRepository: repository,
		passwordHasher:  hasher,
	}
}

func (a *authServer) Create(ctx context.Context, r *auth_v1.CreateRequest) (*auth_v1.CreateResponse, error) {
	if r.GetPassword() == "" {
		return nil, status.Error(codes.InvalidArgument, "user password empty")
	}

	if r.GetPassword() != r.GetPasswordConfirm() {
		return nil, status.Error(codes.InvalidArgument, "password doesn't equal to password_confirm")
	}

	passwordHash, err := a.passwordHasher.HashAndSalt(r.GetPassword())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to generate password hash")
	}

	id, err := a.usersRepository.Create(ctx, r.Info.GetName(), r.Info.GetEmail(), passwordHash, int(r.Info.GetRole()))
	if err != nil {
		log.Printf("Cannot add new user: %v\n", err)
	}

	return &auth_v1.CreateResponse{
		Id: id,
	}, nil
}

func (a *authServer) Get(ctx context.Context, r *auth_v1.GetRequest) (*auth_v1.GetResponse, error) {
	user, err := a.usersRepository.GetByID(ctx, r.GetId())
	if err != nil {
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	return &auth_v1.GetResponse{
		Id: r.GetId(),
		Info: &auth_v1.UserInfo{
			Name:  user.Name,
			Email: user.Email,
			Role:  auth_v1.UserRole(user.UserRole),
		},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}, nil
}

func (a *authServer) Update(ctx context.Context, r *auth_v1.UpdateRequest) (*empty.Empty, error) {
	updateUserInfo := repository.UpdateUserInfo{}
	updateUserInfo.ID = r.GetInfo().Id

	if r.GetInfo().GetName() != nil {
		updateUserInfo.Name = &r.GetInfo().Name.Value
	}

	if r.GetInfo().GetEmail() != nil {
		updateUserInfo.Email = &r.GetInfo().Email.Value
	}

	if r.GetInfo().GetRole() != auth_v1.UserRole_UNKNOWN {
		role := int(r.GetInfo().GetRole())
		updateUserInfo.UserRole = &role
	}

	if _, err := a.usersRepository.Update(ctx, updateUserInfo); err != nil {
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	return &empty.Empty{}, nil
}

func (a *authServer) Delete(ctx context.Context, r *auth_v1.DeleteRequest) (*empty.Empty, error) {
	if _, err := a.usersRepository.Delete(ctx, r.GetId()); err != nil {
		return nil, status.Error(codes.Internal, "failed to delete user")
	}

	return &empty.Empty{}, nil
}
