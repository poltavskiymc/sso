package app

import (
	grpcapp "github.com/poltavskiymc/sso/internal/app/grpc"
	"log/slog"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func NewApp(log *slog.Logger, grpcPort int, storagePath string, tokenTTL time.Duration) *App {
	// TODO: инициализировать хранилище (storage)

	// TODO: init auth srvice (auth)
	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		grpcApp,
	}
}
