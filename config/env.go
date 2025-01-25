package config

import (
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type Env struct {
	HTTPPort         string
	RootApp          string
	AppName          string
	AppEnv           string
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
	PostgresSSLMode  string
	PostgresTimeZone string
}

var GlobalEnv Env

func loadEnv(rootApp string) {
	envPath := filepath.Join(rootApp, ".env")

	err := godotenv.Load(envPath)
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}

	os.Setenv("APP_PATH", rootApp)
	GlobalEnv.RootApp = rootApp

	var ok bool

	GlobalEnv.HTTPPort, ok = os.LookupEnv("HTTP_PORT")
	if !ok {
		panic("missing HTTP_PORT environment")
	}

	GlobalEnv.AppName, ok = os.LookupEnv("APP_NAME")
	if !ok {
		panic("missing APP_NAME environment")
	}

	GlobalEnv.AppEnv, ok = os.LookupEnv("APP_ENV")
	if !ok {
		panic("missing APP_ENV environment")
	}

	GlobalEnv.PostgresHost, ok = os.LookupEnv("POSTGRES_HOST")
	if !ok {
		panic("missing POSTGRES_HOST environment")
	}

	if port, err := strconv.Atoi(os.Getenv("POSTGRES_PORT")); err != nil {
		panic("missing POSTGRES_PORT environment")
	} else {
		GlobalEnv.PostgresPort = int(port)
	}

	GlobalEnv.PostgresUser, ok = os.LookupEnv("POSTGRES_USER")
	if !ok {
		panic("missing POSTGRES_USER environment")
	}

	GlobalEnv.PostgresPassword, ok = os.LookupEnv("POSTGRES_PASSWORD")
	if !ok {
		panic("missing POSTGRES_PASSWORD environment")
	}

	GlobalEnv.PostgresDBName, ok = os.LookupEnv("POSTGRES_DB_NAME")
	if !ok {
		panic("missing POSTGRES_DB_NAME environment")
	}

	GlobalEnv.PostgresSSLMode, ok = os.LookupEnv("POSTGRES_SSLMODE")
	if !ok {
		panic("missing POSTGRES_SSLMODE environment")
	}
}
