package config

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/joho/godotenv"
	"github.com/pujidjayanto/goginboilerplate/pkg/envloader"
)

type database struct {
	host     string
	port     string
	user     string
	password string
	name     string
	ssl      string
}

type server struct {
	port      string
	env       string
	name      string
	secretKey string
}

type redis struct {
	host         string
	port         string
	user         string
	password     string
	databaseName string
}

type config struct {
	database database
	server   server
	redis    redis
}

var (
	globalConfig *config
	loadOnce     sync.Once
	loadError    error
)

// Initialize ensures the configuration is loaded only once
func Initialize() error {
	loadOnce.Do(func() {
		loadError = loadConfiguration()
	})
	return loadError
}

func loadConfiguration() error {
	envPath, err := envloader.GetEnvPath()
	if err != nil || strings.TrimSpace(envPath) == "" {
		return fmt.Errorf("no .env file found")
	}

	err = godotenv.Load(envPath)
	if err != nil {
		return err
	}

	globalConfig = &config{
		database: database{
			host:     os.Getenv("DB_HOST"),
			port:     os.Getenv("DB_PORT"),
			user:     os.Getenv("DB_USER"),
			password: os.Getenv("DB_PASSWORD"),
			name:     os.Getenv("DB_NAME"),
			ssl:      os.Getenv("DB_SSL_MODE"),
		},
		server: server{
			port:      os.Getenv("SERVER_PORT"),
			env:       os.Getenv("SERVER_ENV"),
			name:      os.Getenv("SERVER_NAME"),
			secretKey: os.Getenv("SERVER_SECRET_KEY"),
		},
		redis: redis{
			host:         os.Getenv("REDIS_HOST"),
			port:         os.Getenv("REDIS_PORT"),
			user:         os.Getenv("REDIS_USER"),
			password:     os.Getenv("REDIS_PASSWORD"),
			databaseName: os.Getenv("REDIS_DB_NAME"),
		},
	}

	return nil
}

// Utility methods
func GetDatabaseDSN() string {
	checkInitialized()
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		globalConfig.database.host,
		globalConfig.database.port,
		globalConfig.database.user,
		globalConfig.database.password,
		globalConfig.database.name,
		globalConfig.database.ssl,
	)
}

func GetPort() string {
	checkInitialized()
	port := globalConfig.server.port
	if port == "" {
		port = "3000"
	}
	return fmt.Sprintf(":%s", port)
}

func GetEnv() string {
	checkInitialized()
	return globalConfig.server.env
}

func GetSecretKey() string {
	checkInitialized()
	return globalConfig.server.secretKey
}

func checkInitialized() {
	if globalConfig == nil {
		panic("configuration not initialized. Call Initialize() first")
	}
}
