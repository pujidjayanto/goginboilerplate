package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/joho/godotenv"
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

type Environment struct {
	Database Database
	Server   Server
	Redis    Redis
}

func loadEnvironment() (*Environment, error) {
	envPath, err := getEnvPath()
	if err != nil || strings.TrimSpace(envPath) == "" {
		return nil, fmt.Errorf("no .env file found")
	}

	err = godotenv.Load(envPath)
	if err != nil {
		return nil, err
	}

	return &Environment{
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
	}, nil
}

func getEnvPath() (string, error) {
	directory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filepath := searchup(directory, ".env")
	return filepath, nil
}

func searchup(dir string, filename string) string {
	if dir == "/" || dir == "" || dir == "." {
		return ""
	}

	if _, err := os.Stat(path.Join(dir, filename)); err == nil {
		return path.Join(dir, filename)
	}

	return searchup(path.Dir(dir), filename)
}
