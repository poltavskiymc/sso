package main

import (
	"github.com/poltavskiymc/sso/internal/app"
	"github.com/poltavskiymc/sso/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("config", cfg))

	application := app.NewApp(log, cfg.GRPC.Port, cfg.StoragePath, cfg.TokenTTL)

	// Запускаем сервер в отдельной горутине
	go application.GRPCSrv.MustRun()
	// TODO: инициализировать приложение (app). Точка входа.

	// TODO: запустить gRPC - сервер приложения

	// Создаём слушатель сигнала от системы
	stop := make(chan os.Signal, 1)

	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop
	// Если сигнал поступил, операция перестаёт блокироваться, и идёт дальше

	log.Info("stopping application", slog.String("signal", sign.String()))

	application.GRPCSrv.Stop()

	log.Info("application stopped")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
