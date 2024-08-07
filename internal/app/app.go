package app

import (
	"context"
	"log"
	"net"

	"github.com/andredubov/auth/internal/config"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"github.com/andredubov/golibs/pkg/closer"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// App ...
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
}

// NewApp creates an instance of App struct
func NewApp(ctx context.Context) (*App, error) {
	application := &App{}

	if err := application.initDeps(ctx); err != nil {
		return nil, err
	}

	return application, nil
}

// Run the application
func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	return a.runGRPCServer()
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
	}

	for _, f := range inits {
		if err := f(ctx); err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()

	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	opts := []grpc.ServerOption{
		grpc.Creds(insecure.NewCredentials()),
	}

	a.grpcServer = grpc.NewServer(opts...)
	reflection.Register(a.grpcServer)
	auth_v1.RegisterAuthServer(a.grpcServer, a.serviceProvider.ServerImplementation(ctx))

	return nil
}

func (a *App) runGRPCServer() error {
	log.Printf("GRPC server is running on %s", a.serviceProvider.GRPCConfig().Address())

	listener, err := net.Listen("tcp", a.serviceProvider.GRPCConfig().Address())
	if err != nil {
		return err
	}

	if err = a.grpcServer.Serve(listener); err != nil {
		return err
	}

	return nil
}
