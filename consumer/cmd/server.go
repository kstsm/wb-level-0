package cmd

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gookit/slog"
	"github.com/kstsm/wb-level-0/consumer/config"
	"github.com/kstsm/wb-level-0/consumer/database"
	"github.com/kstsm/wb-level-0/consumer/internal/cache"
	"github.com/kstsm/wb-level-0/consumer/internal/handler"
	"github.com/kstsm/wb-level-0/consumer/internal/mq"
	"github.com/kstsm/wb-level-0/consumer/internal/repository"
	"github.com/kstsm/wb-level-0/consumer/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run() {
	cfg := config.GetConfig()
	validate := validator.New()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	conn := database.InitPostgres(ctx)
	defer conn.Close()

	repo := repository.NewRepository(conn)
	redisClient := cache.NewRedis(ctx, cfg)
	svc := service.NewService(repo, validate, redisClient)

	if err := svc.PreloadCache(ctx); err != nil {
		slog.Error("Failed to preload cache", "error", err)
	}

	router := handler.NewHandler(svc)

	consumer, err := mq.NewConsumer(cfg, svc)
	if err != nil {
		slog.Fatal("Failed to init Kafka consumer", "error", err)
	}

	go consumer.Start(ctx)

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler: router.NewRouter(),
	}

	errChan := make(chan error, 1)

	go func() {
		slog.Infof("Starting server on %s:%d", cfg.Server.Host, cfg.Server.Port)
		errChan <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		slog.Info("Finishing the server...")
	case err = <-errChan:
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Fatal("Error starting server", "error", err)
		}
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		slog.Error("Error while shutting down the server", "error", err)
	}
}
