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

	waitForShutdown(ctx, app)
}

func waitForShutdown(ctx context.Context, app *app.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	sig := <-quit
	log.Printf("Received signal: %v. Initiating shutdown...\n", sig)

	done := make(chan struct{})
	go func() {
		defer close(done)
		app.Shutdown(ctx)
	}()

	select {
	case <-done:
		log.Println("Shutdown completed gracefully.")
	case <-time.After(10 * time.Second):
		log.Println("Forced shutdown due to timeout.")
	}
}
