package main

import (
	"os"
	"os/signal"
	"syscall"

	"service/internal/config"

	"go.uber.org/zap"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	zlog := setupLogger(cfg.Env)
	defer zlog.Sync()
	zlog.Info("logger is configured")

	gracefulShutdown(zlog)
}

func setupLogger(env string) *zap.Logger {
	var logger *zap.Logger

	switch env {
	case envLocal:
		logger = zap.Must(zap.NewDevelopment())
	case envDev:
		logger = zap.Must(zap.NewDevelopment())
	case envProd:
		logger = zap.Must(zap.NewProduction())
	}

	return logger
}

func gracefulShutdown(zlog *zap.Logger) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	s := <-stop

	zlog.Info("Stopping service...", zap.String("signal", s.String()))

	///

	zlog.Info("application stoped")
}
