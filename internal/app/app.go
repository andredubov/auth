package app

import (
	"context"
	"log"
	"net"
	"net/http"
	"sync"

	"github.com/andredubov/auth/internal/interceptor"
	auth_v1 "github.com/andredubov/auth/pkg/auth/v1"
	"github.com/andredubov/golibs/pkg/closer"
	"github.com/andredubov/golibs/pkg/config"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

// App ...
type App struct {
	serviceProvider *serviceProvider
	grpcServer      *grpc.Server
	httpServer      *http.Server
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

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		if err := a.runGRPCServer(); err != nil {
			log.Fatalf("failed to run GRPC server: %v", err)
		}
	}()

	go func() {
		defer wg.Done()

		if err := a.runHTTPServer(); err != nil {
			log.Fatalf("failed to run HTTP server: %v", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initGRPCServer,
		a.initHTTPServer,
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
		grpc.UnaryInterceptor(interceptor.ValidateInterceptor),
	}

	a.grpcServer = grpc.NewServer(opts...)
	reflection.Register(a.grpcServer)
	auth_v1.RegisterAuthServer(a.grpcServer, a.serviceProvider.ServerImplementation(ctx))

	return nil
}

func (a *App) initHTTPServer(ctx context.Context) error {
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := auth_v1.RegisterAuthHandlerFromEndpoint(ctx, mux, a.serviceProvider.GRPCConfig().Address(), opts)
	if err != nil {
		return err
	}

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	a.httpServer = &http.Server{
		Addr:    a.serviceProvider.HTTPConfig().Address(),
		Handler: corsMiddleware.Handler(mux),
		//ReadHeaderTimeout: a.serviceProvider.HTTPConfig().ReadHeaderTimeout(),
	}

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

func (a *App) runHTTPServer() error {
	log.Printf("HTTP server is running on %s", a.serviceProvider.HTTPConfig().Address())

	if err := a.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
