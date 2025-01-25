package config

import (
	"context"
	"fmt"
	"go-codebase/pkg/database/sql"
	"go-codebase/pkg/logger"
)

type Config struct {
	postgres *sql.SQLDatabase
	log      logger.Logger
	env      *Env
}

func NewConfig(ctx context.Context, rootApp string) *Config {
	loadEnv(rootApp)

	cfgChan := make(chan *Config)

	go func() {
		defer close(cfgChan)

		cfg := &Config{}

		cfg.log = logger.NewLogger()

		postgresConfig := &sql.Config{
			Host:     GlobalEnv.PostgresHost,
			Port:     GlobalEnv.PostgresPort,
			User:     GlobalEnv.PostgresUser,
			Password: GlobalEnv.PostgresPassword,
			DBName:   GlobalEnv.PostgresDBName,
			SSLMode:  GlobalEnv.PostgresSSLMode,
		}
		cfg.postgres = sql.NewSQLDatabase(cfg.log, postgresConfig)

		cfgChan <- cfg
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
	return cfg.log
}

func (cfg *Config) Exit(ctx context.Context) {
	cfg.postgres.Close()
	cfg.log.Fatal("\x1b[33;1mConfig: Success close all connection\x1b[0m", "Config.Exit()", "configexit")
}
