package app

import (
	"context"
	"fmt"
	"github.com/osamikoyo/IM-auth/internal/config"
	"github.com/osamikoyo/IM-auth/internal/data"
	"github.com/osamikoyo/IM-auth/internal/server"
	"github.com/osamikoyo/IM-auth/pkg/loger"
	"github.com/osamikoyo/IM-auth/pkg/pb"
	"google.golang.org/grpc"
	"net"
)

type App struct{
	gRPC *grpc.Server
	loger loger.Logger
	config *config.Config
	server *server.Server
}

func Init() (*App, error) {
	cfg, err := config.Load("config.yml")
	if err != nil{
		return nil, err
	}

	st, err := data.New(cfg)
	if err != nil{
		return nil, err
	}

	gserver := grpc.NewServer()
	return &App{
		config: cfg,
		loger: loger.New(),
		gRPC: gserver,
		server: &server.Server{
			Storage:st,
		},
	}, nil
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		a.loger.Info().Msg("server stopped!")
		a.gRPC.Stop()
	}()

	pb.RegisterAuthServiceServer(a.gRPC, a.server)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", a.config.Hostname, a.config.Port))
	if err != nil {
		return err
	}

	a.loger.Info().Str("addr", lis.Addr().String())
	return a.gRPC.Serve(lis)
}