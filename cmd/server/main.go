package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spiderocious/medcord-backend/internal/app"
	"github.com/spiderocious/medcord-backend/internal/configs"
	"github.com/spiderocious/medcord-backend/internal/deps"
	"github.com/spiderocious/medcord-backend/internal/utils/database"
	"github.com/spiderocious/medcord-backend/internal/utils/logger"
)

func main() {
	cfg := configs.Load()
	log := logger.New(cfg.IsProduction)

	connectCtx, cancelConnect := context.WithTimeout(context.Background(), 10*time.Second)
	db, err := database.Connect(connectCtx, cfg.Database, log)
	cancelConnect()
	if err != nil {
		log.Error("connect db", "err", err)
		os.Exit(1)
	}

	d := deps.Wire(cfg, db, log)
	engine := app.New(cfg, d)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", cfg.Port),
		Handler:           engine,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      30 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		log.Info("server starting", "addr", server.Addr, "env", cfg.Env)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("server listen", "err", err)
			os.Exit(1)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Info("shutdown signal received")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Error("server shutdown", "err", err)
	}
	if err := db.Disconnect(shutdownCtx); err != nil {
		log.Error("db disconnect", "err", err)
	}
	log.Info("shutdown complete")
}
