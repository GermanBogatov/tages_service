package app

import (
	"context"
	"fmt"
	pb_prod_images "github.com/GermanBogatov/tages_contracts/gen/go/tages_service/images/v1"
	"github.com/GermanBogatov/tages_service/app/internal/config"
	"github.com/GermanBogatov/tages_service/app/internal/controller/grpc/v1/image"
	"github.com/GermanBogatov/tages_service/app/internal/domain/image/dao"
	"github.com/GermanBogatov/tages_service/app/internal/domain/image/policy"
	"github.com/GermanBogatov/tages_service/app/internal/domain/image/service"
	"github.com/GermanBogatov/tages_service/app/pkg/client/postgresql"
	"github.com/GermanBogatov/tages_service/app/pkg/logging"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"time"
)

type App struct {
	cfg                *config.Config
	grpcServer         *grpc.Server
	pgClient           *pgxpool.Pool
	imageServiceServer pb_prod_images.TagesServiceServer
}

func NewApp(ctx context.Context, config *config.Config) (App, error) {

	pgConfig := postgresql.NewPgConfig(
		config.PostgreSQL.Username, config.PostgreSQL.Password,
		config.PostgreSQL.Host, config.PostgreSQL.Port, config.PostgreSQL.Database,
	)
	pgClient, err := postgresql.NewClient(ctx, 5, time.Second*5, pgConfig)
	if err != nil {
		logging.GetLogger().Fatal(ctx, err)
	}

	imageStorage := dao.NewImageStorage(pgClient)
	imageService := service.NewImageService(imageStorage)
	imagePolicy := policy.NewImagePolicy(imageService)
	imageServiceServer := image.NewServer(
		imagePolicy,
		pb_prod_images.UnimplementedTagesServiceServer{},
	)

	return App{
		cfg:                config,
		pgClient:           pgClient,
		imageServiceServer: imageServiceServer,
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	grp, ctx := errgroup.WithContext(ctx)
	grp.Go(func() error {
		return a.startGRPC(ctx, a.imageServiceServer)
	})
	return grp.Wait()
}

func (a *App) startGRPC(ctx context.Context, server pb_prod_images.TagesServiceServer) error {
	logger := logging.WithFields(ctx, map[string]interface{}{
		"IP":   a.cfg.GRPC.IP,
		"Port": a.cfg.GRPC.Port,
	})
	logger.Info("gRPC Server initializing")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.cfg.GRPC.IP, a.cfg.GRPC.Port))
	if err != nil {
		logger.WithError(err).Fatal("failed to create listener")
	}

	serverOptions := []grpc.ServerOption{}

	a.grpcServer = grpc.NewServer(serverOptions...)

	pb_prod_images.RegisterTagesServiceServer(a.grpcServer, server)

	reflection.Register(a.grpcServer)

	return a.grpcServer.Serve(listener)
}
