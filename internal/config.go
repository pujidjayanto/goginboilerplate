package internal

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/pujidjayanto/goginboilerplate/pkg/envloader"
)

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Ssl      string
}

type Server struct {
	Port string
	Env  string
	Name string
}

type Redis struct {
	Host         string
	Port         string
	User         string
	Password     string
	DatabaseName string
}

type Config struct {
	Database Database
	Server   Server
	Redis    Redis
}

var GlobalConfig *Config

func LoadConfiguration() error {
	envPath, err := envloader.GetEnvPath()
	if err != nil || strings.TrimSpace(envPath) == "" {
		return fmt.Errorf("no .env file found")
	}

	err = godotenv.Load(envPath)
	if err != nil {
		return err
	}

	GlobalConfig = &Config{
		Database: Database{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			Ssl:      os.Getenv("DB_SSL_MODE"),
		},
		Server: Server{
			Port: os.Getenv("SERVER_PORT"),
			Env:  os.Getenv("SERVER_ENV"),
			Name: os.Getenv("SERVER_NAME"),
		},
		Redis: Redis{
			Host:         os.Getenv("REDIS_HOST"),
			Port:         os.Getenv("REDIS_PORT"),
			User:         os.Getenv("REDIS_USER"),
			Password:     os.Getenv("REDIS_PASSWORD"),
			DatabaseName: os.Getenv("REDIS_DB_NAME"),
		},
	}

	return nil
}

func (e *Config) DatabaseDSN() string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		e.Database.Host,
		e.Database.Port,
		e.Database.User,
		e.Database.Password,
		e.Database.Name,
		e.Database.Ssl,
	)
}

func (e *Config) ServerPort() string {
	port := e.Server.Port
	if port == "" {
		port = "3000"
	}

	return fmt.Sprintf(":%s", port)
}
