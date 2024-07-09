package server

import (
	"context"
	"log"

	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"github.com/brianvoe/gofakeit"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type authServer struct {
	auth_v1.UnimplementedAuthServer
}

func NewAuthServer() auth_v1.AuthServer {
	return &authServer{}
}

func (a *authServer) Create(ctx context.Context, r *auth_v1.CreateRequest) (*auth_v1.CreateResponse, error) {

	log.Printf("input = %+v", r)

	return &auth_v1.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func (a *authServer) Get(ctx context.Context, r *auth_v1.GetRequest) (*auth_v1.GetResponse, error) {

	log.Printf("user id = %d", r.GetId())

	return &auth_v1.GetResponse{
		Id: r.GetId(),
		Info: &auth_v1.UserInfo{
			Name:  gofakeit.Name(),
			Email: gofakeit.Email(),
		},
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}, nil
}

func (a *authServer) Update(ctx context.Context, r *auth_v1.UpdateRequest) (*empty.Empty, error) {

	log.Printf("user id: %d", r.GetId())

	return &empty.Empty{}, nil
}

func (a *authServer) Delete(ctx context.Context, r *auth_v1.DeleteRequest) (*empty.Empty, error) {

	log.Printf("user id: %d", r.GetId())

	return &empty.Empty{}, nil
}
