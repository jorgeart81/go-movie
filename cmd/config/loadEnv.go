package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiPort          int
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	DbPort           int
	DbHost           string
	SslMode          string
	Timezone         string
	ConnectTimeout   int
	MdbUrl           string
	MdbApiKey        string
}

func LoadEnv() *Config {
	// load environment variables file .env
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error loading file .env")
	}

	envVars := []string{
		"API_PORT",
		"POSTGRES_USER",
		"POSTGRES_PASSWORD",
		"POSTGRES_DB",
		"DB_PORT",
		"DB_HOST",
		"SSLMODE",
		"TIMEZONE",
		"CONNECT_TIMEOUT",
		"MDB_URL",
		"MDB_API_KEY",
	}

	// Verify that environment variables are defined
	for _, v := range envVars {
		value := os.Getenv(v)
		if value == "" {
			log.Fatalf("Error: environment variable %s is not defined", v)
		}
	}

	apiPort, _ := strconv.Atoi(os.Getenv("API_PORT"))
	postgresUser := os.Getenv("POSTGRES_USER")
	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	postgresDB := os.Getenv("POSTGRES_DB")
	dbPort, _ := strconv.Atoi(os.Getenv("DB_PORT"))
	dbHost := os.Getenv("DB_HOST")
	sslMode := os.Getenv("SSLMODE")
	timezone := os.Getenv("TIMEZONE")
	connectTimeout, _ := strconv.Atoi(os.Getenv("CONNECT_TIMEOUT"))
	mdbUrl := os.Getenv("MDB_URL")
	mdbApiKey := os.Getenv("MDB_API_KEY")

	return &Config{
		ApiPort:          apiPort,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
		PostgresDB:       postgresDB,
		DbPort:           dbPort,
		DbHost:           dbHost,
		SslMode:          sslMode,
		Timezone:         timezone,
		ConnectTimeout:   connectTimeout,
		MdbUrl:           mdbUrl,
		MdbApiKey:        mdbApiKey,
	}
}
