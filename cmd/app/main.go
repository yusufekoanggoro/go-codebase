package main

import (
	"context"
	"fmt"
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
	rootApp, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting working directory: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	defer cancel()

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Failed to start service: %v\nStack trace:\n%s\n", r, debug.Stack())
		}
	}()

	cfg := config.NewConfig(ctx, rootApp)
	defer cfg.Exit(ctx)

	app := app.NewApp(cfg)

	go app.ServeHTTP() // Serve HTTP server in a goroutine

	waitForShutdown(ctx, app)
}

func waitForShutdown(ctx context.Context, app *app.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	sig := <-quit
	fmt.Printf("\nReceived signal: %v. Initiating shutdown...\n", sig)

	app.Shutdown(ctx)
}
