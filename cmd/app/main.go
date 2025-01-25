package main

import (
	"context"
	"go-codebase/config"
	"go-codebase/internal/app"
	"log"
	"os"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Failed to start service: %v\nStack trace:\n%s\n", r, debug.Stack())
		}
	}()

	log.Println("Starting application...")

	cfg := config.NewConfig(ctx)
	defer cfg.Exit(ctx)

	app := app.NewApp(cfg)
	go app.ServeHTTP() // Serve HTTP server in a goroutine

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	sig := <-quit
	log.Printf("Received signal: %v. Initiating shutdown...\n", sig)
	app.Shutdown(ctx)
}
