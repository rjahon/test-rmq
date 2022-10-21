package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	// Environment string
	// ServiceName string

	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	HttpPort string
}

func Load(filenames ...string) Config {
	if err := godotenv.Load(filenames...); err != nil {
		fmt.Println("No .env file found")
	}
	config := Config{}

	// config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "develop"))
	// config.ServiceName = cast.ToString(getOrReturnDefault("SERVICE_NAME", "storage"))

	config.PostgresHost = cast.ToString(getOrReturnDefault("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefault("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefault("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefault("POSTGRES_PASSWORD", "postgres"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefault("POSTGRES_DATABASE", "test"))

	config.HttpPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8008"))
	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	_, exists := os.LookupEnv(key)

	if exists {
		return os.Getenv(key)
	}

	return defaultValue
}
