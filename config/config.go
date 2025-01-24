package config

import (
	"context"
	"fmt"
	"go-codebase/pkg/database/sql"
	"go-codebase/pkg/logger"
)

type Config struct {
	postgres *sql.SQLDatabase
	logger   logger.Logger
	env      *Env
}

func NewConfig(ctx context.Context, rootApp string) *Config {
	loadEnv(rootApp)

	cfgChan := make(chan *Config)

	go func() {
		defer close(cfgChan)

		var cfg Config

		cfg.logger = logger.NewLogger()

		postgresConfig := &sql.Config{
			Host:     GlobalEnv.PostgresHost,
			Port:     GlobalEnv.PostgresPort,
			User:     GlobalEnv.PostgresUser,
			Password: GlobalEnv.PostgresPassword,
			DBName:   GlobalEnv.PostgresDBName,
			SSLMode:  GlobalEnv.PostgresSSLMode,
		}
		cfg.postgres = sql.NewSQLDatabase(cfg.logger, postgresConfig)

		cfgChan <- &cfg
	}()

	// with timeout to init configuration
	select {
	case cfg := <-cfgChan:
		return cfg
	case <-ctx.Done():
		panic(fmt.Errorf("failed to init configuration: %v", ctx.Err()))
	}
}

func (cfg *Config) GetPostgres() *sql.SQLDatabase {
	return cfg.postgres
}

func (cfg *Config) GetLogger() logger.Logger {
	return cfg.logger
}

func (cfg *Config) Exit(ctx context.Context) {
	cfg.postgres.Close()
	cfg.logger.Info("\x1b[33;1mConfig: Success close all connection\x1b[0m", "Config.Exit()", "configexit")
}
