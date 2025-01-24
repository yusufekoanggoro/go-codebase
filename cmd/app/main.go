package main

import (
	"context"
	"fmt"
	"go-codebase/config"
	"go-codebase/internal/app"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"runtime/debug"
	"strings"
	"syscall"
	"time"
)

func main() {
	rootApp, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	rootApp = strings.TrimSuffix(rootApp, "/cmd/app")

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
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
