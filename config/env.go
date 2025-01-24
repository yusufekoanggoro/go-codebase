package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	Port             string
	RootApp          string
	ServiceName      string
	AppEnv           string
	PostgresHost     string
	PostgresPort     string
	PostgresUser     string
	PostgresPassword string
	PostgresDBName   string
	PostgresSSLMode  string
	PostgresTimeZone string
}

var GlobalEnv *Env

func loadEnv(rootApp string) {
	err := godotenv.Load(rootApp + "/.env")
	if err != nil {
		log.Printf("Error loading .env file: %v\n", err)
	}

	os.Setenv("APP_PATH", rootApp)
	GlobalEnv.RootApp = rootApp

	var ok bool

	GlobalEnv.Port, ok = os.LookupEnv("PORT")
	if !ok {
		panic("missing PORT environment")
	}

	GlobalEnv.ServiceName, ok = os.LookupEnv("SERVICE_NAME")
	if !ok {
		panic("missing SERVICE_NAME environment")
	}

	GlobalEnv.AppEnv, ok = os.LookupEnv("APP_ENV")
	if !ok {
		panic("missing APP_ENV environment")
	}

	GlobalEnv.AppEnv, ok = os.LookupEnv("PostgresHost")
	if !ok {
		panic("missing PostgresHost environment")
	}
}
